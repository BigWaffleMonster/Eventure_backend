package api

import (
	"fmt"

	docs "github.com/BigWaffleMonster/Eventure_backend/api/docs"
	"github.com/BigWaffleMonster/Eventure_backend/config"
)

func SwaggerInfo(){
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "Eventura app"
	docs.SwaggerInfo.Description = "Simple app to plan your celebration"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%d", config.GetAppPort())
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}