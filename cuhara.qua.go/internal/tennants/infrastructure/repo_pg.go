package infrastructure

import (
	"context"

	"cuhara.qua.go/internal/tennants/domain"
	"gorm.io/gorm"
)

type PGTennantRepository struct {
	write *gorm.DB
}

func NewPGTennantRepository(write *gorm.DB) *PGTennantRepository {
	return &PGTennantRepository{write: write}
}

func (r *PGTennantRepository) Create(ctx context.Context, t *domain.Tennant, outEvenrs []domain.Event) (int64, error) {
	tx := r.write.WithContext(ctx).Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}

	if err := tx.Create(t).Error; err != nil {
		tx.Rollback()
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		return 0, err
	}

	return t.ID, nil
}
