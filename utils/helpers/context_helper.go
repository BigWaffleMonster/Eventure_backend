package helpers

import (
	"context"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Ключи для хранения в контексте
type contextKey string

const (
	userIDKey      contextKey = "user_id"
	ipKey          contextKey = "ip"
	userAgentKey   contextKey = "user_agent"
	fingerprintKey contextKey = "fingerprint"
)

func FromGinContext(c *gin.Context) context.Context {
	return c.Request.Context()
}

func SetUserID(c *gin.Context, userID uuid.UUID) {
	AddToContext(c, userIDKey, userID)

	c.Set(string(userIDKey), userID)
}

func SetIP(c *gin.Context, ip string) {
	AddToContext(c, ipKey, ip)

	c.Set(string(ipKey), ip)
}

func SetUserAgent(c *gin.Context, userAgent string) {
	AddToContext(c, userAgentKey, userAgent)

	c.Set(string(userAgentKey), userAgent)
}

func SetFingerprint(c *gin.Context, fingerprint string) {
	AddToContext(c, fingerprintKey, fingerprint)

	c.Set(string(fingerprintKey), fingerprint)
}

func GetUserID(ctx context.Context) (uuid.UUID, error) {
	val := ctx.Value(userIDKey)
	if val == nil {
		return uuid.Nil, errors.New("user ID not found in context")
	}

	userID, ok := val.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("invalid user ID type in context")
	}

	return userID, nil
}

func GetIP(ctx context.Context) (string, error) {
	val := ctx.Value(ipKey)
	if val == nil {
		return "", errors.New("IP not found in context")
	}

	ip, ok := val.(string)
	if !ok {
		return "", errors.New("invalid IP type in context")
	}

	return ip, nil
}

func GetUserAgent(ctx context.Context) (string, error) {
	val := ctx.Value(userAgentKey)
	if val == nil {
		return "", errors.New("user agent not found in context")
	}

	userAgent, ok := val.(string)
	if !ok {
		return "", errors.New("invalid user agent type in context")
	}

	return userAgent, nil
}

func GetFingerprint(ctx context.Context) (string, error) {
	val := ctx.Value(fingerprintKey)
	if val == nil {
		return "", errors.New("fingerprint not found in context")
	}

	fingerprint, ok := val.(string)
	if !ok {
		return "", errors.New("invalid fingerprint type in context")
	}

	return fingerprint, nil
}

func AddToContext(c *gin.Context, key interface{}, val interface{}) {
	currentCtx := c.Request.Context()
	currentCtx = context.WithValue(currentCtx, key, val)
	c.Request = c.Request.WithContext(currentCtx)
}