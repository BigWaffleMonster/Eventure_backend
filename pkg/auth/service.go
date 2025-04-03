package auth

import (
	"errors"
	"log"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(data RegisterInput) (string, error)
	Login(data LoginInput) (map[string]string, error)
	RefreshToken(data RefreshInput) (map[string]string, error)
}

type authService struct {
	Repo user.UserRepository
}

func NewAuthService(repo user.UserRepository) AuthService {
	return &authService{Repo: repo}
}

func (s *authService) Register(data RegisterInput) (string, error) {
	var userModel user.User

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

func (s *authService) Login(data LoginInput) (map[string]string, error) {
	if !helpers.IsValidEmail(data.Email) {
		return nil, errors.New("email not valid")
	}

	existingUser, _ := s.Repo.FindByEmail(data.Email)
	if existingUser == nil {
		return nil, errors.New("user doesn`t exists")
	}

	if data.Password != nil {
		passwordHashCheckResult := helpers.CheckPasswordHash(existingUser.Password, *data.Password)
		if !passwordHashCheckResult {
			return nil, errors.New("password don`t match")
		}
	} else {
		//checkTempPassword
	}

	accessToken, err := GenerateAccessToken(existingUser.Email, existingUser.ID)
	if err != nil {
		return nil, errors.New("error Generating Token")
	}

	refreshToken, err := GenerateRefreshToken(existingUser.Email, existingUser.ID)
	if err != nil {
		return nil, errors.New("error Generating Token")
	}

	return map[string]string{"accessToken": accessToken, "refreshToken": refreshToken}, nil
}

func (s *authService) RefreshToken(data RefreshInput) (map[string]string, error) {
	claims, err := ValidateRefreshToken(data.RefreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Generate a new access token
	newAccessToken, err := GenerateAccessToken(claims.Email, claims.ID)
	if err != nil {
		return nil, errors.New("error generating access token")
	}

	// Optionally generate a new refresh token
	newRefreshToken, err := GenerateRefreshToken(claims.Email, claims.ID)
	if err != nil {
		return nil, errors.New("error generating refresh token")
	}

	return map[string]string{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	}, nil
}
