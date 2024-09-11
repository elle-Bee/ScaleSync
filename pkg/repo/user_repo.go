package repo

import (
	"ScaleSync/pkg/models"
	"database/sql"
	"errors"
)

// UserRepository defines the repository for user-related database operations
type UserRepository struct {
	DB *sql.DB
}

// Create inserts a new user into the database
func (r *UserRepository) Create(user *models.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	err := r.DB.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

// ReadAll fetches all users from the database
func (r *UserRepository) ReadAll() ([]*models.User, error) {
	rows, err := r.DB.Query(`SELECT id, name, email FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// Read fetches a single user from the database by ID
func (r *UserRepository) Read(userID int) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow(`SELECT id, name, email FROM users WHERE id = $1`, userID).
		Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

// Update modifies an existing user in the database
func (r *UserRepository) Update(user *models.User) error {
	query := `UPDATE users SET name = $1, email = $2 WHERE id = $3`
	_, err := r.DB.Exec(query, user.Name, user.Email, user.ID)
	return err
}

// Delete removes a user from the database by ID
func (r *UserRepository) Delete(userID int) error {
	_, err := r.DB.Exec(`DELETE FROM users WHERE id = $1`, userID)
	return err
}
