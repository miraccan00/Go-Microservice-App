package services

import (
	"github.com/miraccan00/activity-service/models"
	"github.com/miraccan00/activity-service/repositories"
)

type ActivityService struct {
	ActivityRepo repositories.ActivityRepository
}

func NewActivityService(activityRepo repositories.ActivityRepository) *ActivityService {
	return &ActivityService{ActivityRepo: activityRepo}
}

func (service *ActivityService) CreateActivity(activity *models.Activity) error {
	return service.ActivityRepo.CreateActivity(activity)
}

func (service *ActivityService) GetActivitiesByUserID(userID uint) ([]models.Activity, error) {
	return service.ActivityRepo.GetActivitiesByUserID(userID)
}
