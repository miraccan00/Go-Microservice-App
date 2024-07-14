package repositories

import (
	"github.com/miraccan00/activity-service/models"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	CreateActivity(activity *models.Activity) error
	GetActivitiesByUserID(userID uint) ([]models.Activity, error)
}

type activityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) ActivityRepository {
	return &activityRepository{db}
}

func (r *activityRepository) CreateActivity(activity *models.Activity) error {
	return r.db.Create(activity).Error
}

func (r *activityRepository) GetActivitiesByUserID(userID uint) ([]models.Activity, error) {
	var activities []models.Activity
	err := r.db.Where("user_id = ?", userID).Find(&activities).Error
	return activities, err
}
