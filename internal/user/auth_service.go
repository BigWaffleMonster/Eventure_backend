package user

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/utils"
	"github.com/BigWaffleMonster/Eventure_backend/utils/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(data auth.RegisterInput) (results.Result)
	Login(data auth.LoginInput) (map[string]string, results.Result)
	RefreshToken(data auth.RefreshInput) (map[string]string, results.Result)
}

type authService struct {
	Config utils.ServerConfig
	Repo UserRepository
}

func NewAuthService(repo UserRepository, config utils.ServerConfig) AuthService {
	return &authService{
		Repo: repo,
		Config: config,
	}
}

func (s *authService) Register(data auth.RegisterInput) results.Result {
	var userModel User

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

func (s *authService) Login(data auth.LoginInput) (map[string]string, results.Result) {
	if !helpers.IsValidEmail(data.Email) {
		return nil, results.NewBadRequestError("email not valid")
	}

	existingUser, _ := s.Repo.GetByEmail(data.Email)
	if existingUser == nil {
		return nil, results.NewNotFoundError("User")
	}

	existingUser, err := s.Repo.GetByExpression("email = ?" ,data.Email)
	fmt.Println(err)
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

	accessToken, result := auth.GenerateAccessToken(existingUser.Email, existingUser.ID, s.Config)
	if result.IsFailed {
		return nil, result
	}

	refreshToken, result := s.validateAndGenerateNewRefreshToken(existingUser)
	if result.IsFailed {
		return nil, result
	}

	return map[string]string{
		"accessToken":  accessToken,
		"refreshToken": *refreshToken,
	}, results.NewResultOk()
}

func (s *authService) RefreshToken(data auth.RefreshInput) (map[string]string, results.Result) {
	claims, result := auth.ValidateRefreshToken(data.RefreshToken, s.Config)
	if result.IsFailed {
		return nil, result
	}

	newAccessToken, result := auth.GenerateAccessToken(claims.Email, claims.ID, s.Config)
	if result.IsFailed {
		return nil, result
	}

	existingUser, _ := s.Repo.GetByID(claims.ID)
	if existingUser == nil {
		return nil, results.NewNotFoundError("User")
	}

	newRefreshToken, result := s.validateAndGenerateNewRefreshToken(existingUser)
	if result.IsFailed {
		return nil, result
	}

	return map[string]string{
		"accessToken":  newAccessToken,
		"refreshToken": *newRefreshToken,
	}, results.NewResultOk()
}

func (s *authService) validateAndGenerateNewRefreshToken(existingUser *User) (*string, results.Result) {
	refreshToken, result := s.Repo.GetRefreshToken(existingUser.ID)
	if result.IsFailed {
		return nil, result
	}

	if refreshToken == nil {
		return s.generateNewRefreshToken(existingUser)
	}

	_, result = auth.ValidateRefreshToken(refreshToken.RefreshToken, s.Config)
	if result.IsFailed {
		return s.generateNewRefreshToken(existingUser)
	}

	return &refreshToken.RefreshToken, results.NewResultOk()
}

func (s *authService) generateNewRefreshToken(existingUser *User) (*string, results.Result) {
	newRefreshToken, result := auth.GenerateRefreshToken(existingUser.Email, existingUser.ID, s.Config)
	if result.IsFailed {
		return nil, result
	}

	result = s.Repo.SetRefreshToken(existingUser.ID, newRefreshToken)

	if result.IsFailed {
		return nil, result
	}

	return &newRefreshToken, results.NewResultOk()
}