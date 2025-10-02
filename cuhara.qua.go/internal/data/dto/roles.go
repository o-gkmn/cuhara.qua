package dto

type RoleDTO struct {
	ID     int64
	Name   string
}

type CreateRoleRequest struct {
	Name     string
}

type CreateRoleResponse struct {
	ID int64
}

type UpdateRoleRequest struct {
	ID   int64
	Name *string
}

type UpdateRoleResponse struct {
	ID int64
}

type DeleteRoleRequest struct {
	ID int64
}

type DeleteRoleResponse struct {
	ID int64
}
