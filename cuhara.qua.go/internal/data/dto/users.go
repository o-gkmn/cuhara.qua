package dto

type UserDTO struct {
	ID         int64
	Name       string
	Email      string
	VscAccount string
}

type LoginRequest struct {
	Email    string
	Password string
}

type LoginResponse struct {
	Token string
}

type RegisterRequest struct {
	Name       string
	Email      string
	Password   string
	VscAccount string
	RoleID     int64
	TenantID   int64
}

type RegisterResponse struct {
	ID int64
}

type UpdateUserRequest struct {
	ID         int64
	Name       *string
	Email      *string
	VscAccount *string
	RoleID     *int64
}

type UpdateUserResponse struct {
	ID int64
}

type DeleteUserRequest struct {
	ID int64
}

type DeleteUserResponse struct {
	ID int64
}
