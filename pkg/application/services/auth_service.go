package services

import (
	"errors"
	"log"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/application/views"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain/models"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/domain/utils"
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
	existingUser, _ := s.Repo.FindByUsername(data.UserName)
	if existingUser != nil {
		return "", errors.New("Username already exists")
	}

	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		log.Fatal(err)
		return "", errors.New("Error with hashing password")
	}

	user := models.User{
		ID: uuid.New(),
		UserName: data.UserName,
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
