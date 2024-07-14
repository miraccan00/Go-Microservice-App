package handlers

import (
	"github.com/miraccan00/step-service/models"
	"github.com/miraccan00/step-service/services"

	"github.com/gofiber/fiber/v2"
)

type StepHandler struct {
	StepService *services.StepService
}

func NewStepHandler(stepService *services.StepService) *StepHandler {
	return &StepHandler{StepService: stepService}
}

func (h *StepHandler) CreateStep(c *fiber.Ctx) error {
	step := new(models.Step)
	if err := c.BodyParser(step); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if err := h.StepService.CreateStep(step); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create step"})
	}
	return c.JSON(step)
}

func (h *StepHandler) GetStepsByUserID(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	steps, err := h.StepService.GetStepsByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot get steps"})
	}
	return c.JSON(steps)
}
