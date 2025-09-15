package user

import (
	"cuhara.qua.go/internal/readmodel/role"
	"cuhara.qua.go/internal/readmodel/tennant"
)

type UserDTO struct {
	ID         int64              `json:"id"`
	Name       string             `json:"name"`
	Email      string             `json:"email"`
	VscAccount string             `json:"vscAccount"`
	Role       role.RoleDTO       `json:"role"`
	Tennant    tennant.TennantDTO `json:"tennant"`
}
