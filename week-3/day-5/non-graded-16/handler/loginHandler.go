package handler

import (
	"non-graded-16/config"
	"non-graded-16/entity"
	"golang.org/x/crypto/bcrypt"
)

func Login(email, password string) (entity.User, bool) {
	var user entity.User
	row := config.DB.QueryRow("SELECT id, email, first_name, last_name, date_of_birth, password FROM users WHERE email = ?", email)
	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Birth, &user.Password)
	if err != nil {
		return user, false
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, false
	}
	return user, true
}