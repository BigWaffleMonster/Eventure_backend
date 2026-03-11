package auth

import (
	"net/http"
	"time"

	"github.com/BigWaffleMonster/Eventure_backend/internal/db/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) CreateUser(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	existingUser, err := h.service.GetUserByLogin(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "database_error",
			Message: err.Error(),
		})
		return
	}

	if existingUser != nil {
		c.JSON(http.StatusConflict, ErrorResponse{
			Error:   "email_exists",
			Message: "Пользователь с таким email уже зарегистрирован",
		})

		return
	}

	hashedPassword, err := h.service.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "hashing_error",
			Message: "Ошибка хеширования пароля",
		})
		return
	}

	// 4. Создание нового пользователя
	newUser := schema.User{
		ID:               uuid.New(),
		Email:            req.Email,
		Password:         string(hashedPassword),
		DateCreated:      time.Now(),
		IsEmailConfirmed: false,
	}

	h.service.CreateUser(&newUser)

	c.JSON(http.StatusCreated, "User created")
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "validation_error",
			Message: err.Error(),
		})
		return
	}

	user, err := h.service.GetUserByLogin(req.Login)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Неверный логин или пароль",
		})
		return
	}

	if err := h.service.VerifyPassword(user.Password, req.Password); err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "invalid_credentials",
			Message: "Неверный логин или пароль",
		})
		return
	}

	accessToken, refreshToken, err := h.service.GenerateTokens(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "token_generation_error",
			Message: "Ошибка генерации токенов",
		})
		return
	}

	c.SetCookieData(&http.Cookie{
		Name:        "access_token",
		Value:       accessToken,
		Path:        "/",
		Domain:      "",
		MaxAge:      int(h.service.jwtCfg.AccessTokenExp.Seconds()),
		Secure:      false,
		HttpOnly:    true,
		SameSite:    http.SameSiteDefaultMode,
		Partitioned: false,
		// Expires:  "",
	})

	c.SetCookieData(&http.Cookie{
		Name:        "refresh_token",
		Value:       refreshToken,
		Path:        "/",
		Domain:      "",
		MaxAge:      int(h.service.jwtCfg.RefreshTokenExp.Seconds()),
		Secure:      false,
		HttpOnly:    true,
		SameSite:    http.SameSiteDefaultMode,
		Partitioned: false,
		// Expires:  "",
	})

	// 6. Ответ клиенту
	c.JSON(http.StatusOK, TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		UserID:       user.ID.String(),
		Email:        user.Email,
		Login:        user.Login,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Вы успешно вышли"})
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	refreshTokenString, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "refresh_token_missing",
			Message: "Refresh токен не найден",
		})
		return
	}

	// 2. Валидируем refresh токен
	claims, err := h.service.ValidateRefreshToken(refreshTokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "invalid_refresh_token",
			Message: "Неверный или истекший refresh токен",
		})
		return
	}

	// 3. Получаем пользователя из БД (проверка что пользователь существует)
	user, err := h.service.GetUserByID(claims.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error:   "user_not_found",
			Message: "Пользователь не найден",
		})
		return
	}

	newAccessToken, newRefreshToken, err := h.service.GenerateTokens(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error:   "token_generation_error",
			Message: "Ошибка генерации токенов",
		})
		return
	}

	c.SetCookie(
		"access_token",
		newAccessToken,
		int(h.service.jwtCfg.AccessTokenExp.Seconds()),
		"/",
		"",
		false, // secure: true для HTTPS
		true,  // httpOnly
	)

	c.SetCookie(
		"refresh_token",
		newRefreshToken,
		int(h.service.jwtCfg.RefreshTokenExp.Seconds()),
		"/",
		"",
		false, // secure: true для HTTPS
		true,  // httpOnly
	)

	// 6. Ответ клиенту
	c.JSON(http.StatusOK, TokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		UserID:       user.ID.String(),
		Email:        user.Email,
		Login:        user.Login,
	})
}
