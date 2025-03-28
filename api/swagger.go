package api

import docs "github.com/BigWaffleMonster/Eventure_backend/api/docs"

// @contact.name   Daniil, Sergei, Alex
// @contact.email  rachkov.work@gmail.com, Sergei.m.khanlarov@gmail.com, me@justwalsdi.ru
// @securityDefinitions.basic  Bearer
func SwaggerInfo(){
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "Eventura app"
	docs.SwaggerInfo.Description = "Simple app to plan your celebration"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}