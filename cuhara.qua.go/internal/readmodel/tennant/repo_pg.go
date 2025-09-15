package tennant

import "gorm.io/gorm"

type TennantReadRepository struct {
	read *gorm.DB
}

func NewTennantReadRepository(read *gorm.DB) *TennantReadRepository {
	return &TennantReadRepository{read: read}
}
