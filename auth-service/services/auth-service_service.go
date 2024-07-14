package services

import (
	"github.com/miraccan00/auth-service/models"
	"github.com/miraccan00/auth-service/repositories"
	"github.com/miraccan00/auth-service/utils"
)

type AuthService struct {
	UserRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (service *AuthService) Register(user *models.User) error {
	// Check if username is already taken and user name character length should be greater than 3
	if len(user.Username) < 3 {
		return models.ErrUsernameShort
	}

	existingUser, err := service.UserRepo.FindByUsername(user.Username)
	if err == nil && existingUser != nil {
		return models.ErrUsernameTaken
	}

	return service.UserRepo.CreateUser(user)
}

func (service *AuthService) Login(username, password string) (*models.User, error) {
	user, err := service.UserRepo.FindByUsername(username)
	if err != nil || user.Password != password {
		return nil, err
	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return user, nil
}
