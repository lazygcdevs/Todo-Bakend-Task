package main

import (
	"log"
	"os"

	"todo-api/database"
	"todo-api/handlers"
	"todo-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to database
	database.Connect()

	// Setup Gin router
	router := gin.Default()

	// Setup CORS to allow specific origins (required when using credentials)
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000", // React default
		"http://localhost:5173", // Vite default
		"http://localhost:8080", // Same origin
		"http://127.0.0.1:3000",
		"http://127.0.0.1:5173",
		"http://127.0.0.1:8080",
		"https://todo-backend-app-2024.azurewebsites.net", // Azure App Service
		"https://*.azurewebsites.net",                     // All Azure App Service domains
	}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	// Apply authentication middleware to all routes
	router.Use(middleware.AuthMiddleware())

	// API routes
	api := router.Group("/api/v1")
	{
		api.GET("/todos", handlers.GetTodos)
		api.POST("/todos", handlers.CreateTodo)
		api.PUT("/todos/:id", handlers.UpdateTodo)
		api.DELETE("/todos/:id", handlers.DeleteTodo)
	}

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Todo API is running",
		})
	})

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
