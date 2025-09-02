package middlewares

import (
	"github.com/BigWaffleMonster/Eventure_backend/utils/requests"
	"github.com/gin-gonic/gin"
)

func RequestInfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		requestInfo := requests.GetRequestInfo(c)

		c.Set("request_info", requestInfo)
		
		//TODO: пока оставил, потом переиспользовать при включении логгирования
		// // Логируем информацию
		// start := time.Now()
		// c.Next()
		// duration := time.Since(start)
		
		// // Лог запроса с дополнительной информацией
		// logger.Infof("Request: %s %s | IP: %s | UA: %s | Fingerprint: %s | Duration: %v",
		// 	c.Request.Method,
		// 	c.Request.URL.Path,
		// 	ip,
		// 	userAgent,
		// 	fingerprint,
		// 	duration,
		// )
	}
}