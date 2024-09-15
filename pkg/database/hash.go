package database

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err.Error()
	}
	fmt.Println(bcrypt.CompareHashAndPassword(hash, []byte(password)))
	return string(hash)
}
