package models

type User struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateNewUser(id int, name string, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}
