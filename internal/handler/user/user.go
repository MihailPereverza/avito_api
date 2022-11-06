package user

import (
	"avito_api/internal/db/model"
	"avito_api/internal/service/interface"
	"encoding/json"
	"log"
	"net/http"
)

type UserHandler struct {
	userService service_interface.UserServiceInterface
}

func NewUserHandler(userService service_interface.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := uh.userService.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	if _, err = w.Write(resp); err != nil {
		log.Fatalf("Error happened in ResponseWriter Write. Err: %s", err)
	}
	return
}
