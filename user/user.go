package user

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Usuário
type User struct {
	// Ou seja tags alteram o funcionamento de funcs
	// ... que trabalham com estruturas de dados
	Login        string    `json:"user"`
	HashPassword string    `json:"hash"`
	CreatedAt    time.Time `json:"created_at"`
}

// Create new User
func New(userLogin, userPassword string) User {
	psw, err := HashPassword(userPassword)
	if err != nil {
		fmt.Println("Error ao criar usuário", err)
	}

	return User{
		Login:        userLogin,
		HashPassword: psw,
		CreatedAt:    time.Now(),
	}
}

// Hash password genereator
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}

// Validate password with hash
func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
