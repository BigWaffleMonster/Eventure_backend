package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	configs "github.com/BigWaffleMonster/Eventure_backend/internal/configs"
)

type JWTMiddleware struct {
	cfg *configs.JWTConfig
}

func NewJWTMiddleware(cfg *configs.JWTConfig) *JWTMiddleware {
	return &JWTMiddleware{cfg: cfg}
}

// AuthRequired — middleware для защиты роутов
func (m *JWTMiddleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := m.extractToken(c)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token_required"})
			c.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &configs.Claims{}, func(token *jwt.Token) (any, error) {
			return []byte(m.cfg.AccessSecretKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid_token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*configs.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid_claims"})
			c.Abort()
			return
		}

		// userUUID, err := uuid.ParseBytes(claims.UserID)
		// if err != nil {
		// 	return nil, errors.New("invalid userID")
		// }

		// Сохраняем claims в контекст
		c.Set("userID", claims.UserID)
		c.Set("email", claims.Email)
		c.Set("login", claims.Login)

		c.Next()
	}
}

// extractToken — извлечение токена из cookie или заголовка
func (m *JWTMiddleware) extractToken(c *gin.Context) string {
	// 1. Пробуем из cookie
	cookie, err := c.Cookie("access_token")
	if err == nil && cookie != "" {
		return cookie
	}

	// 2. Пробуем из заголовка Authorization: Bearer <token>
	authHeader := c.GetHeader("Authorization")
	_, found := strings.CutPrefix(authHeader, "Bearer ")

	if found {
		return strings.TrimPrefix(authHeader, "Bearer ")
	} else {
		return ""
	}
}
