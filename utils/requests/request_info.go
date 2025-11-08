package requests

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetClientIP(c *gin.Context) string {
	ip := c.Request.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip = c.Request.Header.Get("X-Forwarded-For")
	if ip != "" {
		ips := strings.Split(ip, ",")
		return strings.TrimSpace(ips[0])
	}

	return c.Request.RemoteAddr
}

func GetUserAgent(c *gin.Context) string {
	return c.Request.Header.Get("User-Agent")
}

func GenerateFingerprint(c *gin.Context) string {
	data := GetClientIP(c) + GetUserAgent(c) + c.Request.Header.Get("Accept-Language")
	
	additionalHeaders := []string{
		c.Request.Header.Get("Sec-Ch-Ua"),
		c.Request.Header.Get("Sec-Ch-Ua-Platform"),
		c.Request.Header.Get("Sec-Ch-Ua-Mobile"),
	}
	
	for _, header := range additionalHeaders {
		if header != "" {
			data += header
		}
	}

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}