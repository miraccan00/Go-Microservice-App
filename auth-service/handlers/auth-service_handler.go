package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/miraccan00/auth-service/models"
	"github.com/miraccan00/auth-service/services"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if err := h.AuthService.Register(user); err != nil {
		switch err {
		case models.ErrUsernameShort:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username is too short"})
		case models.ErrUsernameTaken:
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Username is already taken"})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create user"})
		}
	}
	return c.JSON(user)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	login := new(models.User)
	if err := c.BodyParser(login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	user, err := h.AuthService.Login(login.Username, login.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	return c.JSON(user)
}
