package dto

type UserDTO struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	VscAccount string  `json:"vscAccount"`
	RoleDTO    RoleDTO `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	VscAccount string `json:"vscAccount"`
	RoleID     int64  `json:"roleId"`
}

type RegisterResponse struct {
	ID int64 `json:"id"`
}

type UpdateUserRequest struct {
	ID         int64   `json:"id"`
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	VscAccount *string `json:"vscAccount"`
	RoleID     *int64  `json:"roleId"`
}

type UpdateUserResponse struct {
	ID int64 `json:"id"`
}

type DeleteUserRequest struct {
	ID int64 `json:"id"`
}

type DeleteUserResponse struct {
	ID int64 `json:"id"`
}
