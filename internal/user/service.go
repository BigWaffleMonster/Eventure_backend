package user

import (
	"errors"
	"log"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/internal/user/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/google/uuid"
)

type AuthService struct {
	Repo *UserRepository
}

func NewAuthService(repo *UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(data UserRegisterInput) (string, error) {
	var userModel User

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

func (s *AuthService) Login(data UserLoginInput) (string, error) {
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
	token, err := auth.GenerateAccessToken(existingUser.Email, existingUser.ID)
	if err != nil {
		return "", errors.New("error Generating Token")
	}
	//!TODO add refresh token

	return token, nil
}

func (s *AuthService) RefreshToken() {

}
