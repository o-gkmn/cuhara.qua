package role

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/config"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/models"
	"cuhara.qua.go/internal/util"
	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
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
	log := util.LogFromContext(ctx).With().Str("function", "GetRoles").Logger()
	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return nil, err
	}

	roles, err := models.Roles(
		qm.Load(models.RoleRels.Tenant),
		models.RoleWhere.TenantID.EQ(tenantID),
	).All(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get roles")
		return nil, err
	}

	roleDTOs := make([]dto.RoleDTO, len(roles))
	for i, role := range roles {
		roleDTOs[i] = dto.RoleDTO{
			ID:   role.ID,
			Name: role.Name,
		}
	}

	return roleDTOs, nil
}

func (s *Service) Create(ctx context.Context, request dto.CreateRoleRequest) (dto.CreateRoleResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "CreateRole").Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.CreateRoleResponse{}, err
	}

	exists, err := models.Roles(
		models.RoleWhere.Name.EQ(request.Name),
		models.RoleWhere.TenantID.EQ(tenantID),
	).Exists(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check whether role exists")
		return dto.CreateRoleResponse{}, err
	}

	if exists {
		log.Debug().Str("name", request.Name).Msg("Role already exists")
		return dto.CreateRoleResponse{}, httperrors.ErrConflictRoleAlreadyExists
	}

	role := models.Role{
		Name:      request.Name,
		TenantID:  tenantID,
		CreatedAt: time.Now().UTC(),
	}

	err = role.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("Failed to create role")
		return dto.CreateRoleResponse{}, err
	}

	return dto.CreateRoleResponse{ID: role.ID}, nil
}

func (s *Service) Update(ctx context.Context, request dto.UpdateRoleRequest) (dto.UpdateRoleResponse, error) {
	log := util.LogFromContext(ctx).With().Int64("id", request.ID).Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.UpdateRoleResponse{}, err
	}

	role, err := models.Roles(
		models.RoleWhere.ID.EQ(request.ID),
		models.RoleWhere.TenantID.EQ(tenantID),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Msg("Role not found")
			return dto.UpdateRoleResponse{}, httperrors.ErrRoleNotFound
		}

		log.Error().Err(err).Msg("Failed to fetch role")
		return dto.UpdateRoleResponse{}, err
	}

	changed := false
	if request.Name != nil && role.Name != *request.Name {
		log.Debug().Str("name", *request.Name).Msg("Updating name")

		exists, err := models.Roles(
			models.RoleWhere.Name.EQ(*request.Name),
			models.RoleWhere.TenantID.EQ(role.TenantID),
			models.RoleWhere.ID.NEQ(request.ID),
		).Exists(ctx, s.db)
		if err != nil {
			log.Error().Err(err).Msg("Failed to check whether role exists")
			return dto.UpdateRoleResponse{}, err
		}

		if exists {
			log.Debug().Str("name", *request.Name).Msg("Role already exists")
			return dto.UpdateRoleResponse{}, httperrors.ErrConflictRoleAlreadyExists
		}

		role.Name = *request.Name
		changed = true
	}

	if !changed {
		return dto.UpdateRoleResponse{ID: role.ID}, nil
	}

	role.UpdatedAt = null.TimeFrom(time.Now().UTC())
	_, err = role.Update(ctx, s.db, boil.Whitelist(
		models.RoleColumns.Name,
		models.RoleColumns.UpdatedAt,
	))
	if err != nil {
		log.Error().Err(err).Msg("Failed to update role")
		return dto.UpdateRoleResponse{}, err
	}

	return dto.UpdateRoleResponse{ID: role.ID}, nil
}

func (s *Service) Delete(ctx context.Context, request dto.DeleteRoleRequest) (dto.DeleteRoleResponse, error) {
	log := util.LogFromContext(ctx).With().Int64("id", request.ID).Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.DeleteRoleResponse{}, err
	}

	n, err := models.Roles(
		models.RoleWhere.ID.EQ(request.ID),
		models.RoleWhere.TenantID.EQ(tenantID),
	).DeleteAll(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete role")
		return dto.DeleteRoleResponse{}, err
	}

	if n == 0 {
		return dto.DeleteRoleResponse{}, httperrors.ErrRoleNotFound
	}

	return dto.DeleteRoleResponse(request), nil
}
