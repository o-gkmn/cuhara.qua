package role

import "gorm.io/gorm"

type RoleReadRepository struct {
	read *gorm.DB
}

func NewRoleReadRepository(read *gorm.DB) *RoleReadRepository {
	return &RoleReadRepository{read: read}
}
