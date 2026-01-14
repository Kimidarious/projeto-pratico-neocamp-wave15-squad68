package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68/socialMeli/internal/database"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68/socialMeli/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68/socialMeli/internal/middleware"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68/socialMeli/internal/database"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68/socialMeli/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68/socialMeli/internal/middleware"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(" .env file not found, using environment variables")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf(" Failed to connect to database: %v", err)
	}
	defer database.Close()

	log.Println("Running migrations...")
	if err := database.AutoMigrate(
		db,
		&domain.User{},
		&domain.Follow{},
		&domain.Product{},
		&domain.Post{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Migrations completed")

	r := gin.Default()

	r.Use(middleware.CORS())

	v1 := r.Group("api/v1")
	{
		v1.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"status": "ok"})
		})
	}

	port := "8080"
	log.Printf(" Server running on http:localhost:%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}