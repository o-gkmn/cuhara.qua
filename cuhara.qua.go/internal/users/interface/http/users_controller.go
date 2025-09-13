package http

import (
	"encoding/json"
	"net/http"

	"cuhara.qua.go/internal/common/cqrs"
	usercmd "cuhara.qua.go/internal/users/application/command"
)

type UsersController struct {
	CommandBus *cqrs.CommandBus
}

func NewUsersController(cb *cqrs.CommandBus) *UsersController {
	return &UsersController{CommandBus: cb}
}

// CreateUser godoc
// @Summary      Create a new user
// // @Description  Create a new user with the provided information
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body  usercmd.CreateUserRequest  true " "
// @Success      201   {object}  usercmd.CreateUserResponse
// // @Failure      400   {object}  map[string]string
// @Router       /users [post]
func (uc *UsersController) CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var request usercmd.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
	}

	cmd := usercmd.NewCreateUserCommand(request)

	res, err := cqrs.Execute[usercmd.CreateUserResponse](uc.CommandBus, r.Context(), cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}
