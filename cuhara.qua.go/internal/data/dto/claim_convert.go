package dto

import "cuhara.qua.go/internal/types"

func (c *ClaimDTO) ToTypes() *types.ClaimResponse {
	return &types.ClaimResponse{
		ID: c.ID,
		Name: c.Name,
		Description: c.Description,
	}
}

func (c *CreateClaimResponse) ToTypes() *types.CreateClaimResponse {
	return &types.CreateClaimResponse{
		ID: c.ID,
	}
}

func (c *UpdateClaimResponse) ToTypes() *types.UpdateClaimResponse {
	return &types.UpdateClaimResponse{
		ID: c.ID,
	}
}

func (c *DeleteClaimResponse) ToTypes() *types.DeleteClaimResponse {
	return &types.DeleteClaimResponse{
		ID: c.ID,
	}
}