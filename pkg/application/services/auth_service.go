package services

import (
	"errors"
	"log"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/application/views"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain/models"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/infrastructure/repositories"
	"github.com/google/uuid"
)

type AuthService struct {
	Repo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(data views.UserInfo) (string, error) {
	if !helpers.IsValidEmail(data.Email) {
		return "", errors.New("email not valid")
	}

	existingUser, _ := s.Repo.FindByEmail(data.Email)
	if existingUser != nil {
		return "", errors.New("username already exists")
	}

	hashedPassword, err := helpers.HashPassword(data.Password)
	if err != nil {
		log.Fatal(err)
		return "", errors.New("error with hashing password")
	}

	user := models.User{
		ID: uuid.New(),
		UserName: "",
		Email: data.Email,
		DateCreated: time.Now(),
		Password: hashedPassword,
		IsEmailConfirmed: false,
	}

	err = s.Repo.Create(&user)
	if err != nil {
		return "", err
	}

	return "Successfully created!", nil
}

func (s *AuthService) Login(data views.LoginInfo) (string, error) {
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

	//if no password send email with temp password. Create func validateTempPassword and after it create jwt. if user has not temp password create jwt

	return "Successfully created!", nil
}