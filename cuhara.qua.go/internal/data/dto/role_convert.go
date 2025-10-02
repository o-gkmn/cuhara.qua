package dto

import "cuhara.qua.go/internal/types"

func (r RoleDTO) ToTypes() *types.RoleResponse {
	return &types.RoleResponse{
		Id:     &r.ID,
		Name:   &r.Name,
	}
}

func (c CreateRoleResponse) ToTypes() *types.CreateRoleResponse {
	return &types.CreateRoleResponse{
		Id: &c.ID,
	}
}

func (r UpdateRoleResponse) ToTypes() *types.UpdateRoleResponse {
	return &types.UpdateRoleResponse{
		Id: &r.ID,
	}
}

func (d DeleteRoleResponse) ToTypes() *types.DeleteRoleResponse {
	return &types.DeleteRoleResponse{
		Id: &d.ID,
	}
}