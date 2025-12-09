package user

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Login     string    `json:"user"`
	Password  string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
}

// Create new User
func New(userLogin, userPassword string) User {
	psw, err := HashPassword(userPassword)
	if err != nil {
		fmt.Println("Error ao criar usuário", err)
	}

	return User{
		Login:     userLogin,
		Password:  psw,
		CreatedAt: time.Now(),
	}
}

// Hasd password genereator
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
