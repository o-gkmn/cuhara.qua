package dto

import "cuhara.qua.go/internal/types"

func (t *TenantDTO) ToTypes() *types.TenantResponse {
	return &types.TenantResponse{
		ID: t.ID,
		Name: t.Name,
	}
}

func (c *CreateTenantResponse) ToTypes() *types.CreateTenantResponse {
	return &types.CreateTenantResponse{
		ID: c.ID,
	}
}

func (u *UpdateTenantResponse) ToTypes() *types.UpdateTenantResponse {
	return &types.UpdateTenantResponse{
		ID: u.ID,
	}
}

func (d *DeleteTenantResponse) ToTypes() *types.DeleteTenantResponse {
	return &types.DeleteTenantResponse{
		ID: d.ID,
	}
}