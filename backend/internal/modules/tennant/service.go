package tenant

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
	"cuhara.qua.go/internal/util/db"
	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
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

func (s *Service) GetAll(ctx context.Context) ([]dto.TenantDTO, error) {
	log := util.LogFromContext(ctx).With().Str("function", "GetAll").Logger()

	tenants, err := models.Tenants().All(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenants")
		return nil, err
	}

	tenantDTOs := make([]dto.TenantDTO, len(tenants))
	for i, tenant := range tenants {
		tenantDTOs[i] = dto.TenantDTO{
			ID:   tenant.ID,
			Name: tenant.Name,
		}
	}

	log.Debug().Msg("Tenants fetched successfully")

	return tenantDTOs, nil
}

func (s *Service) Create(ctx context.Context, request dto.CreateTenantRequest) (dto.CreateTenantResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "Create").Logger()

	//Checked tenant existence
	exists, err := models.Tenants(
		models.TenantWhere.Name.EQ(request.Name),
	).Exists(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check whether tenant exist")
		return dto.CreateTenantResponse{}, err
	}

	if exists {
		log.Debug().Str("name", request.Name).Msg("Tenant already exists")
		return dto.CreateTenantResponse{}, httperrors.ErrConflitTenantAlreadyExist
	}

	//started transaction. if any error appears rolled back
	var result dto.CreateTenantResponse
	err = db.WithTransaction(ctx, s.db, func(ce boil.ContextExecutor) error {
		tenant := &models.Tenant{
			Name:      request.Name,
			CreatedAt: time.Now().UTC(),
		}

		if err := tenant.Insert(ctx, ce, boil.Infer()); err != nil {
			log.Err(err).Msg("Failed to insert tenant")
			return err
		}

		result.ID = tenant.ID

		return nil
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to run transaction")
		return dto.CreateTenantResponse{}, err
	}

	log.Debug().Msg("Tenant created successfully")

	return result, nil
}

func (s *Service) Update(ctx context.Context, request dto.UpdateTenantRequest) (dto.UpdateTenantResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "Update").Logger()

	t, err := models.Tenants(
		models.TenantWhere.ID.EQ(request.ID),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Msg("Tenant not found")
			return dto.UpdateTenantResponse{}, httperrors.ErrTenantNotFound
		}

		log.Error().Err(err).Msg("Failed to fetch tenant")
		return dto.UpdateTenantResponse{}, err
	}

	changed := false
	if request.Name != nil && t.Name != *request.Name {
		log.Debug().Str("name", *request.Name).Msg("Updating name")

		//Check name uniqueness
		exists, err := models.Tenants(
			models.TenantWhere.Name.EQ(*request.Name),
			models.TenantWhere.ID.NEQ(request.ID),
		).Exists(ctx, s.db)
		if err != nil {
			log.Error().Err(err).Str("name", *request.Name).Msg("Failed to check name uniqueness")
			return dto.UpdateTenantResponse{}, err
		}

		if exists {
			log.Debug().Str("name", *request.Name).Msg("Tenant name already in use")
			return dto.UpdateTenantResponse{}, httperrors.ErrConflitTenantAlreadyExist
		}

		t.Name = *request.Name
		changed = true
	}

	if !changed {
		return dto.UpdateTenantResponse{ID: t.ID}, nil
	}

	t.UpdatedAt = null.TimeFrom(time.Now().UTC())
	_, err = t.Update(ctx, s.db, boil.Whitelist(
		models.TenantColumns.Name,
		models.TenantColumns.UpdatedAt,
	))
	if err != nil {
		log.Error().Err(err).Msg("Failed to update tenant")
		return dto.UpdateTenantResponse{}, err
	}

	log.Debug().Msg("Tenant updated successfully")

	return dto.UpdateTenantResponse{ID: t.ID}, nil
}

func (s *Service) Delete(ctx context.Context, request dto.DeleteTenantRequest) (dto.DeleteTenantResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "Delete").Logger()

	tenant, err := models.Tenants(
		models.TenantWhere.ID.EQ(request.ID),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Msg("Tenant not found")
			return dto.DeleteTenantResponse{}, httperrors.ErrTenantNotFound
		}

		log.Error().Err(err).Msg("Failed to fetch tenant")
		return dto.DeleteTenantResponse{}, err
	}

	_, err = tenant.Delete(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete tenant")
		return dto.DeleteTenantResponse{}, err
	}

	log.Debug().Msg("Tenant deleted successfully")

	return dto.DeleteTenantResponse{ID: tenant.ID}, nil
}
