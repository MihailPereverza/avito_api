package handler

import (
	"avito_api/internal/service/interface"
)

type Handler struct {
	userService *service_interface.UserServiceInterface
}
