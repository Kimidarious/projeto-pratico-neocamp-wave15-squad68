package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/database"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/domain"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/handler"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/middleware"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/repository"
	"github.com/Kimidarious/projeto-pratico-neocamp-wave15-squad68.git/internal/service"
)


func main() {
	loadEnv()

	
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}
	defer database.Close()

	
	log.Println("🔄 Running migrations...")
	if err := database.AutoMigrate(
		db,
		&domain.User{},
		&domain.Follow{},
		&domain.Product{},
		&domain.Post{},
	); err != nil {
		log.Fatalf("❌ Failed to migrate database: %v", err)
	}
	log.Println("✅ Migrations completed")


	userRepo := repository.NewUserRepository(db)
	
	userService := service.NewUserService(userRepo)
	
	userCRUDHandler := handler.NewUserCRUDHandler(userService)
		
	r := gin.Default()
	r.Use(middleware.CORS())


	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status":  "ok",
			"message": "API is running",
		})
	})
	
	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.POST("", userCRUDHandler.CreateUser)
			users.GET("", userCRUDHandler.GetAllUsers)
			users.GET("/:id", userCRUDHandler.GetUser)
			users.PUT("/:id", userCRUDHandler.UpdateUser)
			users.DELETE("/:id", userCRUDHandler.DeleteUser)
		}
	}


	port := "8080"
	log.Println("===========================================")
	log.Printf("🚀 Server running on http://localhost:%s", port)
	log.Printf("📋 Health Check: http://localhost:%s/health", port)
	log.Printf("👤 Users API: http://localhost:%s/api/v1/users", port)
	log.Println("===========================================")
	
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}

func loadEnv() {
	if err := godotenv.Load(); err == nil {
		log.Println("✅ Loaded .env from current directory")
		return
	} else {
		log.Printf("⚠️  Could not load .env from current directory: %v", err)
	}


	wd, err := os.Getwd()
	if err != nil {
		log.Printf("⚠️  Failed to get working directory: %v", err)
		log.Println("⚠️  .env file not found, using environment variables")
		return
	}

	dir := wd
	for {
		envPath := filepath.Join(dir, ".env")
		if _, err := os.Stat(envPath); err == nil {
			loadErr := godotenv.Load(envPath)
			if loadErr == nil {
				log.Printf("✅ Loaded .env from %s", envPath)
				return
			}
			log.Printf("⚠️  Failed to load .env from %s: %v", envPath, loadErr)
			return
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	log.Println("⚠️  .env file not found, using environment variables")
}