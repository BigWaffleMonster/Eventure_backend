package service

import (
	"errors"
	"log"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/models"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/repository"
	"github.com/BigWaffleMonster/Eventure_backend/pkg/utils"
)

type AuthService struct {
	Repo *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(data models.User) (string, error) {
	existingUser, _ := s.Repo.FindByUsername(data.UserName)
	if existingUser != nil {
		return "", errors.New("Username already exists")
	}

	data.DateCreated = time.Now()

	hashedPassword, err := utils.HashPassword(data.Password)
	if err != nil {
		log.Fatal(err)
		return "", errors.New("Error with hashing password")
	}

	data.Password = hashedPassword

	err = s.Repo.Create(&data)
	if err != nil {
		return "", err
	}

	return "Successfully created!", nil
}
