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
	log.Println("✅ Migrations completed successfully")

	userRepo := repository.NewUserRepository(db)
	followRepo := repository.NewFollowRepository(db)
	productRepo := repository.NewProductRepository(db)
	postRepo := repository.NewPostRepository(db)

	userService := service.NewUserService(userRepo)
	followService := service.NewFollowService(followRepo, userRepo)
	postService := service.NewPostService(postRepo, productRepo, userRepo, followRepo)

	userCRUDHandler := handler.NewUserCRUDHandler(userService)
	userHandler := handler.NewUserHandler(followService)
	productHandler := handler.NewProductHandler(postService)

	r := gin.Default()

	r.Use(middleware.CORS())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "SocialMeli API is running",
			"version": "1.0",
		})
	})

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{

			users.POST("/:id/follow/:userIdToFollow", userHandler.FollowUser)

			users.POST("/:id/unfollow/:userIdToFollow", userHandler.UnfollowUser)

			users.GET("/:id/followers/count", userHandler.GetFollowersCount)

			users.GET("/:id/followers/list", userHandler.GetFollowersList)

			users.GET(":id/followed/list", userHandler.GetFollowedList)

			users.POST("", userCRUDHandler.CreateUser)
			users.GET("", userCRUDHandler.GetAllUsers)
			users.GET("/:id", userCRUDHandler.GetUser)      
			users.PUT("/:id", userCRUDHandler.UpdateUser)    
			users.DELETE("/:id", userCRUDHandler.DeleteUser)   
		}

		products := v1.Group("/products")
		{
			products.POST("/post", productHandler.CreatePost)

			products.POST("/promo-post", productHandler.CreatePromoPost)

			products.GET("/followed/:id/list", productHandler.GetFollowedPosts)

			products.GET("/:id/countPromo", productHandler.CountPromoProducts)
		}
	}

	port := getEnv("SERVER_PORT", "8080")
	
	log.Println("=================================================")
	log.Printf("🚀 SocialMeli API Server")
	log.Println("=================================================")
	log.Printf("📡 Server:    http://localhost:%s", port)
	log.Printf("🏥 Health:    http://localhost:%s/health", port)
	log.Printf("📚 API v1:    http://localhost:%s/api/v1", port)
	log.Println("-------------------------------------------------")
	log.Printf("👤 Users:     http://localhost:%s/api/v1/users", port)
	log.Printf("📦 Products:  http://localhost:%s/api/v1/products", port)
	log.Println("=================================================")
	log.Println("✅ Server is ready to accept connections")
	log.Println("🔄 Press CTRL+C to stop")
	log.Println("=================================================")

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

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}