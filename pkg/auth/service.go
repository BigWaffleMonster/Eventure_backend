package auth

import (
	"log"
	"strings"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/internal/user"
	"github.com/BigWaffleMonster/Eventure_backend/utils"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(data RegisterInput) (results.Result)
	Login(data LoginInput) (map[string]string, results.Result)
	RefreshToken(data RefreshInput) (map[string]string, results.Result)
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

func (s *authService) Register(data RegisterInput) results.Result {
	var userModel user.User

	if !helpers.IsValidEmail(data.Email) {
		return results.NewBadRequestError("email not valid")
	}

	existingUser, _ := s.Repo.GetByEmail(data.Email)
	if existingUser != nil {
		return results.NewConflictError("email already exists")
	}

	if data.Password != nil {
		hashedPassword, err := helpers.HashPassword(*data.Password)
		if err != nil {
			log.Fatal(err)
			return results.NewInternalError("Error with hashing password")
		}

		userModel.Password = hashedPassword
	}

	userModel.ID = uuid.New()
	userModel.Email = data.Email
	userModel.DateCreated = time.Now()
	userModel.UserName = strings.Split(data.Email,`@`)[0]

	return s.Repo.Create(&userModel)
}

func (s *authService) Login(data LoginInput) (map[string]string, results.Result) {
	if !helpers.IsValidEmail(data.Email) {
		return nil, results.NewBadRequestError("email not valid")
	}

	existingUser, _ := s.Repo.GetByEmail(data.Email)
	if existingUser == nil {
		return nil, results.NewNotFoundError("User")
	}

	if data.Password != nil {
		passwordHashCheckResult := helpers.CheckPasswordHash(*data.Password, existingUser.Password)
		if !passwordHashCheckResult {
			return nil, results.NewUnauthorizedError("Password or login is incorrect")
		}
	} else {
		//TODO: checkTempPassword
	}

	accessToken, result := GenerateAccessToken(existingUser.Email, existingUser.ID, s.Config)
	if result.IsFailed {
		return nil, result
	}

	refreshToken, result := GenerateRefreshToken(existingUser.Email, existingUser.ID, s.Config)

	if result.IsFailed {
		return nil, result
	}

	result = s.Repo.SetRefreshToken(existingUser.ID, refreshToken)

	if result.IsFailed {
		return nil, result
	}

	return map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}, results.NewResultOk()
}

func (s *authService) RefreshToken(data RefreshInput) (map[string]string, results.Result) {
	claims, result := ValidateRefreshToken(data.RefreshToken, s.Config)
	if result.IsFailed {
		return nil, result
	}

	result = s.Repo.GetRefreshToken(data.RefreshToken)
	if result.IsFailed {
		return nil, result
	}

	newAccessToken, result := GenerateAccessToken(claims.Email, claims.ID, s.Config)
	if result.IsFailed {
		return nil, result
	}

	newRefreshToken, result := GenerateRefreshToken(claims.Email, claims.ID, s.Config)
	if result.IsFailed {
		return nil, result
	}

	return map[string]string{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	}, results.NewResultOk()
}
