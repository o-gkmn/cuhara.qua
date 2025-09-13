package infrastructure

import (
	"context"

	"cuhara.qua.go/internal/roles/domain"
	"gorm.io/gorm"
)

type PGRoleRepository struct {
	write *gorm.DB
}

func NewPGRoleRepository(write *gorm.DB) *PGRoleRepository {
	return &PGRoleRepository{write: write}
}

func (r *PGRoleRepository) Create(ctx context.Context, role *domain.Role, outEvents []domain.Event) (int64, error) {
	tx := r.write.WithContext(ctx).Begin()
	if tx.Error != nil {
		return 0, nil
	}

	if err := tx.Create(role).Error; err != nil {
		tx.Rollback()
		return 0, nil
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return role.ID, nil
}
