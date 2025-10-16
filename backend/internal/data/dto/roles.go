package dto

type RoleDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateRoleRequest struct {
	Name string `json:"name"`
}

type CreateRoleResponse struct {
	ID int64 `json:"id"`
}

type UpdateRoleRequest struct {
	ID   int64   `json:"id"`
	Name *string `json:"name"`
}

type UpdateRoleResponse struct {
	ID int64 `json:"id"`
}

type DeleteRoleRequest struct {
	ID int64 `json:"id"`
}

type DeleteRoleResponse struct {
	ID int64 `json:"id"`
}
