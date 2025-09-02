package requests

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/gin-gonic/gin"
)

type RequestInfo struct{
	IP string
	UserAgent string
	Fingerprint string
}

func GetClientIP(c *gin.Context) string {
	// Проверяем различные заголовки, которые могут содержать реальный IP
	// (актуально при работе behind reverse proxy like nginx)
	ip := c.Request.Header.Get("X-Real-IP")
	if ip != "" {
		return ip
	}

	ip = c.Request.Header.Get("X-Forwarded-For")
	if ip != "" {
		// X-Forwarded-For может содержать список IP через запятую
		ips := strings.Split(ip, ",")
		return strings.TrimSpace(ips[0])
	}

	// Если заголовков нет, используем RemoteAddr
	return strings.Split(c.Request.RemoteAddr, ":")[0]
}

func GetUserAgent(c *gin.Context) string {
	return c.Request.Header.Get("User-Agent")
}

func GenerateFingerprint(c *gin.Context) string {
	// Используем комбинацию данных для создания уникального отпечатка
	data := GetClientIP(c) + GetUserAgent(c) + c.Request.Header.Get("Accept-Language")
	
	// Добавляем дополнительные заголовки для уникальности
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

	// Хэшируем данные для получения фиксированной длины
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func GetRequestInfo(c *gin.Context) (RequestInfo) {
	return RequestInfo{
		IP: GetClientIP(c),
		UserAgent: GetUserAgent(c),
		Fingerprint: GenerateFingerprint(c),
	}
}