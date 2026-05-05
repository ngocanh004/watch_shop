package main

import (
	"log"
	"os"

	"watch-shop/backend/internal/config"
	"watch-shop/backend/internal/database"
	"watch-shop/backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	gin.SetMode(cfg.GinMode)

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Khong the ket noi database: %v", err)
	}
	if err := database.Migrate(db); err != nil {
		log.Fatalf("Migrate that bai: %v", err)
	}

	if len(os.Args) > 1 && os.Args[1] == "seed" {
		database.Seed(db)
		return
	}

	database.Seed(db)

	r := routes.Setup(db, cfg)
	log.Printf("Server chay tai cong %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
