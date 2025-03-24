package auth_service

import (
	"errors"
	"log"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/models"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/models/http_models"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/service/auth_service/helpers"
	"github.com/google/uuid"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(data http_models.UserRegisterInput) (string, error) {
	var userModel models.User

	if !helpers.IsValidEmail(data.Email) {
		return "", errors.New("email not valid")
	}

	existingUser, _ := s.Repo.FindByEmail(data.Email)
	if existingUser != nil {
		return "", errors.New("email already exists")
	}

	if data.Password != nil {
		hashedPassword, err := helpers.HashPassword(*data.Password)
		if err != nil {
			log.Fatal(err)
			return "", errors.New("Error with hashing password")
		}

		userModel.Password = hashedPassword
	}

	userModel.ID = uuid.New()
	userModel.Email = data.Email
	userModel.DateCreated = time.Now()

	err := s.Repo.Create(&userModel)
	if err != nil {
		return "", err
	}

	return "Successfully created!", nil
}

func (s *AuthService) Login(data http_models.UserLoginInput) (string, error) {
	if !helpers.IsValidEmail(data.Email) {
		return "", errors.New("email not valid")
	}

	existingUser, _ := s.Repo.FindByEmail(data.Email)
	if existingUser == nil {
		return "", errors.New("user doesn`t exists")
	}

	if data.Password != nil {
		passwordHashCheckResult := helpers.CheckPasswordHash(existingUser.Password, *data.Password)
		if !passwordHashCheckResult {
			return "", errors.New("password don`t match")
		}
	} else {
		//checkTempPassword
	}
	//create jwt and send to user
	token, err := helpers.GenerateAccessToken(existingUser.Email, existingUser.ID)
	if err != nil {
		return "", errors.New("error Generating Token")
	}
	//!TODO add refresh token

	return token, nil
}

func (s *AuthService) RefreshToken() {

}
