package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/miraccan00/apigateway-service/middleware"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	// Public routes for auth service
	app.Post("/register", middleware.ProxyRequest("http://auth-service:3001/register"))
	app.Post("/login", middleware.ProxyRequest("http://auth-service:3001/login"))

	// Protected routes for activity and step services
	app.Use(middleware.TokenValidator)

	app.Post("/api/activities", middleware.ProxyRequest("http://activity-service:3002/api/activities"))
	app.Get("/api/activities/:userID", middleware.ProxyRequest("http://activity-service:3002/api/activities/:userID"))

	app.Post("/api/steps", middleware.ProxyRequest("http://step-service:3003/api/steps"))
	app.Get("/api/steps/:userID", middleware.ProxyRequest("http://step-service:3003/api/steps/:userID"))

	app.Listen(":3000")
}
