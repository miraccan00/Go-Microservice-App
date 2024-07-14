package handlers

import (
	"github.com/miraccan00/activity-service/models"
	"github.com/miraccan00/activity-service/services"

	"github.com/gofiber/fiber/v2"
)

type ActivityHandler struct {
	ActivityService *services.ActivityService
}

func NewActivityHandler(activityService *services.ActivityService) *ActivityHandler {
	return &ActivityHandler{ActivityService: activityService}
}

func (h *ActivityHandler) CreateActivity(c *fiber.Ctx) error {
	activity := new(models.Activity)
	if err := c.BodyParser(activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if err := h.ActivityService.CreateActivity(activity); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot create activity"})
	}
	return c.JSON(activity)
}

func (h *ActivityHandler) GetActivitiesByUserID(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("userID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	activities, err := h.ActivityService.GetActivitiesByUserID(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Cannot get activities"})
	}
	return c.JSON(activities)
}
