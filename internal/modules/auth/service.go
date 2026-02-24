package auth

import (
	"errors"
	"time"

	configs "github.com/BigWaffleMonster/Eventure_backend/internal/configs"
	schema "github.com/BigWaffleMonster/Eventure_backend/internal/db/schema"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo   *AuthRepository
	jwtCfg *configs.JWTConfig
}

func NewAuthService(repo *AuthRepository, jwtCfg *configs.JWTConfig) *AuthService {
	return &AuthService{repo: repo, jwtCfg: jwtCfg}
}

func (s *AuthService) GetUserByID(id uuid.UUID) (*schema.User, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) GetUserByLogin(login string) (*schema.User, error) {
	user, err := s.repo.GetUserByLogin(login)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) CreateUser(newUser *schema.User) error {
	if err := s.repo.CreateUser(newUser); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, ErrorResponse{
		// 	Error:   "hashing_error",
		// 	Message: "Ошибка хеширования пароля",
		// })
		return nil, err
	}

	return hashedPassword, nil
}

func (s *AuthService) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (s *AuthService) GenerateTokens(user *schema.User) (string, string, error) {
	expiresAt := time.Now().Add(s.jwtCfg.AccessTokenExp)

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, configs.Claims{
		UserID: user.ID,
		Email:  user.Email,
		Login:  user.Login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.jwtCfg.Issuer,
			Subject:   user.ID.String(),
		},
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.jwtCfg.AccessSecretKey))
	if err != nil {
		return "", "", err
	}

	refreshExpiresAt := time.Now().Add(s.jwtCfg.RefreshTokenExp)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, configs.Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    s.jwtCfg.Issuer,
			Subject:   user.ID.String(),
		},
	})

	refreshTokenString, err := refreshToken.SignedString([]byte(s.jwtCfg.RefreshSecretKey))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *AuthService) ValidateRefreshToken(tokenString string) (*configs.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &configs.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.jwtCfg.RefreshSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*configs.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("неверные claims токена")
	}

	if time.Until(claims.ExpiresAt.Time) < s.jwtCfg.AccessTokenExp {
		return nil, errors.New("это не refresh токен")
	}

	return claims, nil
}
