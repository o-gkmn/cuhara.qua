package dto

import "cuhara.qua.go/internal/types"

func (r RoleDTO) ToTypes() *types.RoleResponse {
	return &types.RoleResponse{
		ID:     r.ID,
		Name:   r.Name,
		Tenant: r.Tenant.ToTypes(),
	}
}