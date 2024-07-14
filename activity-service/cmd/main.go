package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/miraccan00/activity-service/handlers"
	"github.com/miraccan00/activity-service/middleware"
	"github.com/miraccan00/activity-service/models"
	"github.com/miraccan00/activity-service/repositories"
	"github.com/miraccan00/activity-service/services"
)

var dsn = "host=db user=postgres password=postgres dbname=activitydb port=5432 sslmode=disable"

func main() {
	// Database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&models.Activity{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories and services
	activityRepo := repositories.NewActivityRepository(db)
	activityService := services.NewActivityService(activityRepo)
	activityHandler := handlers.NewActivityHandler(activityService)

	app := fiber.New()
	app.Use(logger.New())

	// JWT Middleware for protected routes
	app.Use(middleware.JWTProtected())

	// Protected routes
	app.Post("/api/activities", activityHandler.CreateActivity)
	app.Get("/api/activities/:userID", activityHandler.GetActivitiesByUserID)

	log.Fatal(app.Listen(":3002"))
}
