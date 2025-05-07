package auth

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/BigWaffleMonster/Eventure_backend/utils"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(data RegisterInput) (string, error)
	Login(data LoginInput) (map[string]string, error)
	RefreshToken(data RefreshInput) (map[string]string, error)
}

type authService struct {
	Config utils.ServerConfig
	Repo user.UserRepository
}

func NewAuthService(repo user.UserRepository, config utils.ServerConfig) AuthService {
	return &authService{
		Repo: repo,
		Config: config,
	}
}

func (s *authService) Register(data RegisterInput) (string, error) {
	var userModel user.User

	if !helpers.IsValidEmail(data.Email) {
		return "", errors.New("email not valid")
	}

	existingUser, _ := s.Repo.GetByEmail(data.Email)
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
	userModel.UserName = strings.Split(data.Email,`@`)[0]

	err := s.Repo.Create(&userModel)
	if err != nil {
		return "", err
	}

	return "Successfully created!", nil
}

func (s *authService) Login(data LoginInput) (map[string]string, error) {
	if !helpers.IsValidEmail(data.Email) {
		return nil, errors.New("email not valid")
	}

	existingUser, _ := s.Repo.GetByEmail(data.Email)
	if existingUser == nil {
		return nil, errors.New("user doesn`t exists")
	}

	if data.Password != nil {
		passwordHashCheckResult := helpers.CheckPasswordHash(*data.Password, existingUser.Password)
		if !passwordHashCheckResult {
			return nil, errors.New("password don`t match")
		}
	} else {
		//checkTempPassword
	}

	accessToken, err := GenerateAccessToken(existingUser.Email, existingUser.ID, s.Config)
	if err != nil {
		return nil, errors.New("error Generating Token")
	}

	refreshToken, err := GenerateRefreshToken(existingUser.Email, existingUser.ID, s.Config)
	if err != nil {
		return nil, errors.New("error Generating Token")
	}

	err = s.Repo.SetRefreshToken(existingUser.ID, refreshToken)
	if err != nil {
		return nil, errors.New("can`t set refresh token")
	}

	return map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}, nil
}

func (s *authService) RefreshToken(data RefreshInput) (map[string]string, error) {
	claims, err := ValidateRefreshToken(data.RefreshToken, s.Config)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	err = s.Repo.GetRefreshToken(data.RefreshToken)
	if err != nil {
		return nil, errors.New("refresh token doesn`t exists")
	}

	// Generate a new access token
	newAccessToken, err := GenerateAccessToken(claims.Email, claims.ID, s.Config)
	if err != nil {
		return nil, errors.New("error generating access token")
	}

	// Optionally generate a new refresh token
	newRefreshToken, err := GenerateRefreshToken(claims.Email, claims.ID, s.Config)
	if err != nil {
		return nil, errors.New("error generating refresh token")
	}

	return map[string]string{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	}, nil
}
