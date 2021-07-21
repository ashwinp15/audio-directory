package model

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/ashwinp15/audio-directory/database"
	"golang.org/x/crypto/bcrypt"
)

func (user *Creator) Create(password string) error {
	id := rand.Intn(1000)
	sql := fmt.Sprintf(`
INSERT INTO creators(id, name, email, password)
	 VALUES ($1, $2, $3, $4)
	 `)
	hashed, err := HashPassword(password)
	if err != nil {
		return err
	}
	database.PGclient.Exec(context.TODO(), sql, id, user.Name, user.Email, hashed)
	return nil
}

func HashPassword(password string) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return bytes, err
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserByEmail(email string) (*Creator, error) {

	sql := fmt.Sprintf(`
SELECT name, email FROM creators
	 WHERE email = $1
	 `)

	var user Creator
	row := database.PGclient.QueryRow(context.TODO(), sql, email)
	if err := row.Scan(&user.Name, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}
