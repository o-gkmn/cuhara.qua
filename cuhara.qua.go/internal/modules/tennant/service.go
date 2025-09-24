package tennant

import (
	"context"
	"database/sql"

	"cuhara.qua.go/internal/config"
	"cuhara.qua.go/internal/data/dto"
)

type Service struct {
	db     *sql.DB
	config config.Server
}

func NewService(config config.Server, db *sql.DB) *Service {
	return &Service{
		config: config,
		db:     db,
	}
}

func (s *Service) GetTennants(ctx context.Context) ([]dto.TennantDTO, error) {
	panic("not implemented")
}

func (s *Service) Create(ctx context.Context, request dto.CreateTennantRequest) (dto.CreateTennantResponse, error) {
	panic("not implemented")
}

func (s *Service) Update(ctx context.Context, request dto.UpdateTennantRequest) (dto.UpdateTennantResponse, error) {
	panic("not implemented")
}

func (s *Service) Delete(ctx context.Context, request dto.DeleteTennantRequest) (dto.DeleteTennantResponse, error) {
	panic("not implemented")
}
