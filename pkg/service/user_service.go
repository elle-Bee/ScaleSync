package service

import (
	"ScaleSync/pkg/models"
	"ScaleSync/pkg/repo"
	"errors"
)

// UserService defines the interface for user-related business operations
type UserService interface {
	CreateUser(user *models.User) error
	GetAllUsers() ([]*models.User, error)
	GetUser(userID int) (*models.User, error)
	UpdateUser(user *models.User) error
	DeleteUser(userID int) error
}

// userServiceImpl is the concrete implementation of the UserService interface
type userServiceImpl struct {
	Repo repo.UserRepository
}

// NewUserService returns an implementation of UserService
func NewUserService(repo repo.UserRepository) UserService {
	return &userServiceImpl{
		Repo: repo,
	}
}

// CreateUser handles business logic and creates a user in the repository
func (s *userServiceImpl) CreateUser(user *models.User) error {
	// Example validation logic
	if user.Name == "" || user.Email == "" {
		return errors.New("invalid user data")
	}
	return s.Repo.Create(user)
}

// GetAllUsers fetches all users from the repository
func (s *userServiceImpl) GetAllUsers() ([]*models.User, error) {
	return s.Repo.ReadAll()
}

// GetUser fetches a single user by ID from the repository
func (s *userServiceImpl) GetUser(userID int) (*models.User, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.Repo.Read(userID)
}

// UpdateUser handles business logic and updates a user in the repository
func (s *userServiceImpl) UpdateUser(user *models.User) error {
	if user.ID <= 0 {
		return errors.New("invalid user ID")
	}
	return s.Repo.Update(user)
}

// DeleteUser deletes a user by ID from the repository
func (s *userServiceImpl) DeleteUser(userID int) error {
	if userID <= 0 {
		return errors.New("invalid user ID")
	}
	return s.Repo.Delete(userID)
}
