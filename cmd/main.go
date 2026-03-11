package main

import (
	"fmt"
	"log"

	api "github.com/BigWaffleMonster/Eventure_backend/internal/api"
	c "github.com/BigWaffleMonster/Eventure_backend/internal/configs"
	database "github.com/BigWaffleMonster/Eventure_backend/internal/db"
)

func main() {
	config, err := c.InitConfig()
	if err != nil {
		log.Fatalf("❌ Init config error: %v", err)
	}

	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("❌ Не удалось подключиться к БД: %v", err)
	}
	log.Println("✅ Подключение к БД установлено")

	router := api.InitRouter(config, db)
	addr := fmt.Sprintf("localhost:%s", config.Server.Port)

	log.Printf("🚀 Сервер запущен на http://localhost%s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("❌ Ошибка запуска сервера: %v", err)
	}
}
