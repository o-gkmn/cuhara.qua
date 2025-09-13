package infrastructure

import (
	"context"

	"gorm.io/gorm"

	"cuhara.qua.go/internal/users/domain"
)

type PGUserRepository struct {
	write *gorm.DB
}

func NewPGUserRepository(write *gorm.DB) *PGUserRepository {
	return &PGUserRepository{write: write}
}

func (r *PGUserRepository) Create(ctx context.Context, u *domain.User, outEvents []domain.Event) (int64, error) {
	tx := r.write.WithContext(ctx).Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	if err := tx.Create(u).Error; err != nil {
		tx.Rollback()
		return 0, nil
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return u.ID, nil
}
