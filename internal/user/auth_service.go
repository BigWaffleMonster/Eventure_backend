package user

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/utils"
	"github.com/BigWaffleMonster/Eventure_backend/utils/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/requests"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(data auth.RegisterInput) (results.Result)
	Login(data auth.LoginInput, requestInfo requests.RequestInfo) (map[string]string, results.Result)
	RefreshToken(data auth.RefreshInput, requestInfo requests.RequestInfo) (map[string]string, results.Result)
	Logout(data auth.RefreshInput, requestInfo requests.RequestInfo) (results.Result)
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

func (s *authService) Login(data auth.LoginInput, requestInfo requests.RequestInfo) (map[string]string, results.Result) {
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

	accessToken, result := auth.GenerateAccessToken(existingUser.Email, existingUser.ID, s.Config)
	if result.IsFailed {
		return nil, result
	}

	refreshToken, result := s.createUserSession(existingUser, requestInfo)
	if result.IsFailed {
		return nil, result
	}

	return map[string]string{
		"accessToken":  accessToken,
		"refreshToken": *refreshToken,
	}, results.NewResultOk()
}

func (s *authService) RefreshToken(data auth.RefreshInput, requestInfo requests.RequestInfo) (map[string]string, results.Result) {
	claims, result := auth.ValidateRefreshToken(data.RefreshToken, s.Config)

	if result.IsFailed {
		return nil, results.NewUnauthorizedError("Invalid token")
	}

	existingUser, _ := s.Repo.GetByID(claims.ID)
	if existingUser == nil {
		return nil, results.NewNotFoundError("User")
	}

	newRefreshToken, result := s.updateUserSession(existingUser, claims.SessionID, data.RefreshToken, requestInfo)
	if result.IsFailed {
		return nil, result
	}

	newAccessToken, result := auth.GenerateAccessToken(existingUser.Email, claims.ID, s.Config)
	if result.IsFailed {
		return nil, result
	}

	return map[string]string{
		"accessToken":  newAccessToken,
		"refreshToken": *newRefreshToken,
	}, results.NewResultOk()
}

func (s *authService) Logout(data auth.RefreshInput, requestInfo requests.RequestInfo) (results.Result) {
	/*
	@author Sergey Khanlarov
	@data 03.09.2025
	Поскольку логаут это удаление сессии, то если мы не находим сессию или не удается 
	найти пользователя и тд и тп, то мы просто возвращаем ок
	*/
	claims, result := auth.ValidateRefreshToken(data.RefreshToken, s.Config)

	if result.IsFailed {
		return results.NewResultOk()
	}

	existingUser, _ := s.Repo.GetByID(claims.ID)
	if existingUser == nil {
		return results.NewResultOk()
	}

	result = s.Repo.DeleteUserSession(claims.SessionID)

	return results.NewResultOk()
}

func (s *authService) updateUserSession(
	existingUser *User, 
	sessionID uuid.UUID, 
	refreshToken string, 
	requestInfo requests.RequestInfo) (*string, results.Result) {

	session, result := s.getUserSession(sessionID)

	fmt.Println(sessionID.String())

	if result.IsFailed {
		return nil, result
	}

	if session.ExpiresAt.Before(time.Now()) {
		return nil, results.NewUnauthorizedError("Authorization failed")
	}

	result = s.Repo.DeleteUserSession(session.ID)

	if result.IsFailed {
		return nil, result
	}

	return s.createUserSession(existingUser, requestInfo)
}

func (s *authService) getUserSession(sessionID uuid.UUID) (*UserSession, results.Result) {
	session, result := s.Repo.GetUserSession(sessionID)

	if result.IsFailed {
		return nil, results.NewUnauthorizedError("Authorization failed")
	}
	
	if session == nil {
		return nil, results.NewUnauthorizedError("Authorization failed")
	}

	return session, results.NewResultOk()
}

func (s *authService) createUserSession(existingUser *User, requestInfo requests.RequestInfo) (*string, results.Result) {
	sessionID, result := s.Repo.CreateUserSession(existingUser.ID, requestInfo)

	if result.IsFailed {
		return nil, results.NewUnauthorizedError("Authorization failed")
	}

	newRefreshToken, result := auth.GenerateRefreshToken(existingUser.ID, *sessionID, s.Config)
	if result.IsFailed {
		return nil, results.NewUnauthorizedError("Authorization failed")
	}

	return &newRefreshToken, results.NewResultOk()
}