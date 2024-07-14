package services

import (
	"github.com/miraccan00/step-service/models"
	"github.com/miraccan00/step-service/repositories"
)

type StepService struct {
	StepRepo repositories.StepRepository
}

func NewStepService(stepRepo repositories.StepRepository) *StepService {
	return &StepService{StepRepo: stepRepo}
}

func (service *StepService) CreateStep(step *models.Step) error {
	return service.StepRepo.CreateStep(step)
}

func (service *StepService) GetStepsByUserID(userID uint) ([]models.Step, error) {
	return service.StepRepo.GetStepsByUserID(userID)
}
