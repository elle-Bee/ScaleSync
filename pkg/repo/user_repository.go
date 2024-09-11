package repo

import (
	db "ScaleSync/pkg/database"
	"ScaleSync/pkg/models"
	"context"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.User, error)
	GetUser(id string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id string) error
	GetAllUsers() ([]models.User, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(user *models.User) (*models.User, error) {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := db.Pool.QueryRow(context.Background(), query, user.Name, user.Email).Scan(&user.ID)
	return user, err
}

func (r *userRepository) GetUser(id string) (*models.User, error) {
	var user models.User
	query := `SELECT id, name, email FROM users WHERE id = $1`
	err := db.Pool.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email)
	return &user, err
}

func (r *userRepository) UpdateUser(user *models.User) (*models.User, error) {
	query := `UPDATE users SET name=$1, email=$2 WHERE id=$3`
	_, err := db.Pool.Exec(context.Background(), query, user.Name, user.Email, user.ID)
	return user, err
}

func (r *userRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id=$1`
	_, err := db.Pool.Exec(context.Background(), query, id)
	return err
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	rows, err := db.Pool.Query(context.Background(), `SELECT id, name, email FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
