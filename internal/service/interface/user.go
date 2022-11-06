package service_interface

import "avito_api/internal/db/model"

type UserServiceInterface interface {
	CreateUser(user *model.User) (*model.User, error)
}
