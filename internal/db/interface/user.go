package db_interface

import "avito_api/internal/db/model"

type UserDB interface {
	CreateUser(user *model.User) (*model.User, error)
}
