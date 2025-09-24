package role

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

func (s *Service) GetRoles(ctx context.Context) ([]dto.RoleDTO, error) {
	panic("not implemented")
}

func (s *Service) Create(ctx context.Context, request dto.CreateRoleRequest) (dto.CreateRoleResponse, error) {
	panic("not implemented")
}

func (s *Service) Update(ctx context.Context, request dto.UpdateRoleRequest) (dto.UpdateRoleResponse, error) {
	panic("not implemented")
}

func (s *Service) Delete(ctx context.Context, request dto.DeleteRoleRequest) (dto.DeleteRoleResponse, error) {
	panic("not implemented")
}
