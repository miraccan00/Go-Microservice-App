package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/miraccan00/step-service/handlers"
	"github.com/miraccan00/step-service/middleware"
	"github.com/miraccan00/step-service/models"
	"github.com/miraccan00/step-service/repositories"
	"github.com/miraccan00/step-service/services"
)

var dsn = "host=db user=postgres password=postgres dbname=stepdb port=5432 sslmode=disable"

func main() {
	// Database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Step{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories and services
	stepRepo := repositories.NewStepRepository(db)
	stepService := services.NewStepService(stepRepo)
	stepHandler := handlers.NewStepHandler(stepService)

	app := fiber.New()
	app.Use(logger.New())

	// JWT Middleware for protected routes
	app.Use(middleware.JWTProtected())

	// Protected routes
	app.Post("/api/steps", stepHandler.CreateStep)
	app.Get("/api/steps/:userID", stepHandler.GetStepsByUserID)

	log.Fatal(app.Listen(":3003"))
}
