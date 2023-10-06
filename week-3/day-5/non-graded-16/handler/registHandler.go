package handler

import (
	"non-graded-16/config"
	"non-graded-16/entity"
	"golang.org/x/crypto/bcrypt"
)

func Register(user entity.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec("INSERT INTO users (email, first_name, last_name, date_of_birth, password) VALUES (?,?,?,?,?)", user.Email, user.FirstName, user.LastName, user.Birth, hashedPassword)
	return err
}