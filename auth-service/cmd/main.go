package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/miraccan00/auth-service/handlers"
	"github.com/miraccan00/auth-service/models"
	"github.com/miraccan00/auth-service/repositories"
	"github.com/miraccan00/auth-service/services"
)

var dsn = "host=db user=postgres password=postgres dbname=authdb port=5432 sslmode=disable"

func main() {
	// Database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories and services
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	app := fiber.New()
	app.Use(logger.New())

	// Public routes
	app.Post("/register", authHandler.Register)
	app.Post("/login", authHandler.Login)

	log.Fatal(app.Listen(":3001"))
}
