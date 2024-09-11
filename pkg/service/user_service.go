package service

import (
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/repo"
)

type UserService interface {
	CreateUser(user UserDTO) (*models.User, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(id string, user UserDTO) (*models.User, error)
	DeleteUser(id string) error
	GetAllUsers() ([]models.User, error)
}

type userService struct {
	UserRepo repo.UserRepository
}

// UserDTO represents data transfer object for user
type UserDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserService(repo repo.UserRepository) UserService {
	return &userService{
		UserRepo: repo,
	}
}

func (s *userService) CreateUser(user UserDTO) (*models.User, error) {
	newUser := &models.User{
		Name:  user.Name,
		Email: user.Email,
	}

	createdUser, err := s.UserRepo.CreateUser(newUser)
	return createdUser, err
}

func (s *userService) GetUser(id string) (*models.User, error) {
	return s.UserRepo.GetUser(id)
}

func (s *userService) UpdateUser(id string, user UserDTO) (*models.User, error) {
	existingUser, err := s.UserRepo.GetUser(id)
	if err != nil {
		return nil, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email

	return s.UserRepo.UpdateUser(existingUser)
}

func (s *userService) DeleteUser(id string) error {
	return s.UserRepo.DeleteUser(id)
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.GetAllUsers()
}
