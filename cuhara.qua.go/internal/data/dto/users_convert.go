package dto

import "cuhara.qua.go/internal/types"

func (u UserDTO) ToTypes() *types.UserResponse {
	return &types.UserResponse{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		VscAccount: u.VscAccount,
	}
}

func (l LoginResponse) ToTypes() *types.LoginResponse {
	return &types.LoginResponse{
		Token: l.Token,
	}
}

func (r RegisterResponse) ToTypes() *types.RegisterResponse {
	return &types.RegisterResponse{
		ID: r.ID,
	}
}

func (u UpdateUserResponse) ToTypes() *types.UpdateUserResponse {
	return &types.UpdateUserResponse{
		ID: u.ID,
	}
}

func (d DeleteUserResponse) ToTypes() *types.DeleteUserResponse {
	return &types.DeleteUserResponse{
		ID: d.ID,
	}
}
