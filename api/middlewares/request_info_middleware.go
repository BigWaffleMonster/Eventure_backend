package middlewares

import (
	"github.com/BigWaffleMonster/Eventure_backend/utils/helpers"
	"github.com/BigWaffleMonster/Eventure_backend/utils/requests"
	"github.com/gin-gonic/gin"
)

func RequestInfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		helpers.SetIP(c, requests.GetClientIP(c))
		helpers.SetUserAgent(c, requests.GetUserAgent(c))
		helpers.SetFingerprint(c, requests.GenerateFingerprint(c))
	}
}