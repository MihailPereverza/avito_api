package user

import (
	"avito_api/internal/db/interface"
	"avito_api/internal/db/model"
	"database/sql"
	"errors"
)

type DBUser struct {
	db *sql.DB
}

var duplicateNameErrorDB = `ERROR: duplicate key value violates unique constraint "users_name_key" (SQLSTATE 23505)`
var duplicateNameError = errors.New("username is already used")

var duplicateEmailErrorDB = `ERROR: duplicate key value violates unique constraint "users_email_key" (SQLSTATE 23505)`
var duplicateEmailError = errors.New("email is already used")

func NewDBUser(db *sql.DB) db_interface.UserDB {
	return &DBUser{db}
}

func (D DBUser) CreateUser(user *model.User) (*model.User, error) {
	query := `INSERT INTO users(name, email) VALUES ($1, $2) RETURNING user_id`
	err := D.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
	// лучше через свои типы ошибок, но там долго(((
	if err != nil && err.Error() == duplicateNameErrorDB {
		err = duplicateNameError
	} else if err != nil && err.Error() == duplicateEmailErrorDB {
		err = duplicateEmailError
	}
	return user, err
}
