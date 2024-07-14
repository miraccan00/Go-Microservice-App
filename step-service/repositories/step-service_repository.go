package repositories

import (
	"github.com/miraccan00/step-service/models"
	"gorm.io/gorm"
)

type StepRepository interface {
	CreateStep(step *models.Step) error
	GetStepsByUserID(userID uint) ([]models.Step, error)
}

type stepRepository struct {
	db *gorm.DB
}

func NewStepRepository(db *gorm.DB) StepRepository {
	return &stepRepository{db}
}

func (r *stepRepository) CreateStep(step *models.Step) error {
	return r.db.Create(step).Error
}

func (r *stepRepository) GetStepsByUserID(userID uint) ([]models.Step, error) {
	var steps []models.Step
	err := r.db.Where("user_id = ?", userID).Find(&steps).Error
	return steps, err
}
