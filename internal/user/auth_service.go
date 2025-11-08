package user

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/pkg/auth"
	"github.com/BigWaffleMonster/Eventure_backend/utils/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/results"
	"github.com/BigWaffleMonster/Eventure_backend/utils/validators"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(ctx context.Context, data auth.RegisterInput) (results.Result)
	Login(ctx context.Context, data auth.LoginInput) (map[string]string, results.Result)
	RefreshToken(ctx context.Context, data auth.RefreshInput) (map[string]string, results.Result)
	Logout(ctx context.Context, data auth.RefreshInput) (results.Result)
}

type authService struct {
	Repo UserRepository
}

func NewAuthService(repo UserRepository) AuthService {
	return &authService{
		Repo: repo,
	}
}

func (s *authService) Register(ctx context.Context, data auth.RegisterInput) results.Result {
	var userModel User

	if !validators.IsValidEmail(data.Email) {
		return results.NewBadRequestError("email not valid")
	}

	existingUser, _ := s.Repo.GetByEmail(ctx, data.Email)
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

	return s.Repo.Create(ctx, &userModel)
}

func (s *authService) Login(ctx context.Context, data auth.LoginInput) (map[string]string, results.Result) {
	if !validators.IsValidEmail(data.Email) {
		return nil, results.NewBadRequestError("email not valid")
	}

	existingUser, _ := s.Repo.GetByEmail(ctx, data.Email)
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

	accessToken, result := auth.GenerateAccessToken(ctx, existingUser.Email, existingUser.ID)
	if result.IsFailed {
		return nil, result
	}

	refreshToken, result := s.createUserSession(ctx, existingUser)
	if result.IsFailed {
		return nil, result
	}

	return map[string]string{
		"accessToken":  accessToken,
		"refreshToken": *refreshToken,
	}, results.NewResultOk()
}

func (s *authService) RefreshToken(ctx context.Context, data auth.RefreshInput) (map[string]string, results.Result) {
	claims, result := auth.ValidateRefreshToken(ctx, data.RefreshToken)

	if result.IsFailed {
		return nil, results.NewUnauthorizedError("Invalid token")
	}

	existingUser, _ := s.Repo.GetByID(ctx, claims.ID)
	if existingUser == nil {
		return nil, results.NewNotFoundError("User")
	}

	newRefreshToken, result := s.updateUserSession(ctx, existingUser, claims.SessionID, data.RefreshToken)
	if result.IsFailed {
		return nil, result
	}

	newAccessToken, result := auth.GenerateAccessToken(ctx, existingUser.Email, claims.ID)
	if result.IsFailed {
		return nil, result
	}

	return map[string]string{
		"accessToken":  newAccessToken,
		"refreshToken": *newRefreshToken,
	}, results.NewResultOk()
}

func (s *authService) Logout(ctx context.Context, data auth.RefreshInput) (results.Result) {
	/*
	@author Sergey Khanlarov
	@data 03.09.2025
	Поскольку логаут это удаление сессии, то если мы не находим сессию или не удается 
	найти пользователя и тд и тп, то мы просто возвращаем ок
	*/
	claims, result := auth.ValidateRefreshToken(ctx, data.RefreshToken)

	if result.IsFailed {
		return results.NewResultOk()
	}

	existingUser, _ := s.Repo.GetByID(ctx, claims.ID)
	if existingUser == nil {
		return results.NewResultOk()
	}

	result = s.Repo.DeleteUserSession(ctx, claims.SessionID)

	return results.NewResultOk()
}

func (s *authService) updateUserSession(
	ctx context.Context, 
	existingUser *User, 
	sessionID uuid.UUID, 
	refreshToken string) (*string, results.Result) {

	session, result := s.getUserSession(ctx,sessionID)

	fmt.Println(sessionID.String())

	if result.IsFailed {
		return nil, result
	}

	if session.ExpiresAt.Before(time.Now()) {
		return nil, results.NewUnauthorizedError("Authorization failed")
	}

	result = s.Repo.DeleteUserSession(ctx, session.ID)

	if result.IsFailed {
		return nil, result
	}

	return s.createUserSession(ctx, existingUser)
}

func (s *authService) getUserSession(ctx context.Context, sessionID uuid.UUID) (*UserSession, results.Result) {
	session, result := s.Repo.GetUserSession(ctx, sessionID)

	if result.IsFailed {
		return nil, results.NewUnauthorizedError("Authorization failed")
	}
	
	if session == nil {
		return nil, results.NewUnauthorizedError("Authorization failed")
	}

	return session, results.NewResultOk()
}

func (s *authService) createUserSession(ctx context.Context, existingUser *User) (*string, results.Result) {
	sessionID, result := s.Repo.CreateUserSession(ctx, existingUser.ID)

	if result.IsFailed {
		return nil, results.NewUnauthorizedError("Authorization failed")
	}

	newRefreshToken, result := auth.GenerateRefreshToken(ctx, existingUser.ID, *sessionID)
	if result.IsFailed {
		return nil, results.NewUnauthorizedError("Authorization failed")
	}

	return &newRefreshToken, results.NewResultOk()
}