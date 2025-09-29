package dto

import "cuhara.qua.go/internal/types"

func (u UserDTO) ToTypes() *types.UserResponse {
	return &types.UserResponse{
		Id:         &u.ID,
		Name:       &u.Name,
		Email:      &u.Email,
		VscAccount: &u.VscAccount,
	}
}

func (l LoginResponse) ToTypes() *types.LoginResponse {
	return &types.LoginResponse{
		Token: &l.Token,
	}
}

func (r RegisterResponse) ToTypes() *types.RegisterResponse {
	return &types.RegisterResponse{
		Id: &r.ID,
	}
}

func (u UpdateUserResponse) ToTypes() *types.UpdateUserResponse {
	return &types.UpdateUserResponse{
		Id: &u.ID,
	}
}

func (d DeleteUserResponse) ToTypes() *types.DeleteUserResponse {
	return &types.DeleteUserResponse{
		Id: &d.ID,
	}
}
