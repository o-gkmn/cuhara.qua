package user

import (
	"strings"

	"cuhara.qua.go/internal/readmodel/role"
	"cuhara.qua.go/internal/readmodel/tennant"
	"cuhara.qua.go/internal/users/domain"
	"gorm.io/gorm"
)

type UserReadRepository struct {
	read *gorm.DB
}

func NewUserReadRepository(read *gorm.DB) *UserReadRepository {
	return &UserReadRepository{read: read}
}

func (r *UserReadRepository) GetByEmail(email string) (*UserDTO, error) {
	var u domain.User
	if err := r.read.
		Where("email = ?", strings.ToLower(strings.TrimSpace(email))).
		Preload("Role").
		Preload("Tennant").
		First(&u).Error; err != nil {
		return nil, err
	}

	dto := &UserDTO{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		VscAccount: u.VscAccount,
	}

	if u.Role != nil {
		dto.Role = role.RoleDTO{
			ID:   u.Role.ID,
			Name: u.Role.Name,
		}
	}

	if u.Tennant != nil {
		dto.Tennant = tennant.TennantDTO{
			ID:   u.Tennant.ID,
			Name: u.Tennant.Name,
		}
	}

	return dto, nil
}
