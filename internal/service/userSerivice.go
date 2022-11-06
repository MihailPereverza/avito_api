package service

import (
	"avito_api/internal/db/interface"
	"avito_api/internal/db/model"
)

type UserService struct {
	userDB db_interface.UserDB
}

func NewUserService(dbUser db_interface.UserDB) *UserService {
	return &UserService{
		dbUser,
	}
}

func (us *UserService) CreateUser(user *model.User) (*model.User, error) {
	return us.userDB.CreateUser(user)
}
