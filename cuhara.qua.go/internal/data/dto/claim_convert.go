package dto

import "cuhara.qua.go/internal/types"

func (c *ClaimDTO) ToTypes() *types.ClaimResponse {
	return &types.ClaimResponse{
		Id: &c.ID,
		Name: &c.Name,
		Description: &	c.Description,
	}
}

func (c *CreateClaimResponse) ToTypes() *types.CreateClaimResponse {
	return &types.CreateClaimResponse{
		Id: &c.ID,
	}
}

func (c *UpdateClaimResponse) ToTypes() *types.UpdateClaimResponse {
	return &types.UpdateClaimResponse{
		Id: &c.ID,
	}
}

func (c *DeleteClaimResponse) ToTypes() *types.DeleteClaimResponse {
	return &types.DeleteClaimResponse{
		Id: &c.ID,
	}
}