package claim

import (
	"context"
	"database/sql"
	"errors"

	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/config"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/models"
	"cuhara.qua.go/internal/util"
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

func (s *Service) GetAll(ctx context.Context) ([]dto.ClaimDTO, error) {
	log := util.LogFromContext(ctx).With().Str("function", "GetAll").Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return nil, err
	}

	claims, err := models.Claims(
		models.ClaimWhere.TenantID.EQ(tenantID),
	).All(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get claims")
		return nil, err
	}

	claimDTOs := make([]dto.ClaimDTO, len(claims))
	for i, claim := range claims {
		claimDTOs[i] = dto.ClaimDTO{
			ID:          claim.ID,
			Name:        claim.Name,
			Description: claim.Description.String,
		}
	}

	log.Debug().Msg("Claims fetched successfully")

	return claimDTOs, nil
}

func (s *Service) Create(ctx context.Context, request dto.CreateClaimRequest) (dto.CreateClaimResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "Create").Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.CreateClaimResponse{}, err
	}

	exists, err := models.Claims(
		models.ClaimWhere.Name.EQ(request.Name),
		models.ClaimWhere.TenantID.EQ(tenantID),
	).Exists(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to check whether claim exists")
		return dto.CreateClaimResponse{}, err
	}

	if exists {
		log.Debug().Str("name", request.Name).Msg("Claim already exists")
		return dto.CreateClaimResponse{}, httperrors.ErrConflictClaimAlreadyExists
	}

	claim := models.Claim{
		Name:        request.Name,
		Description: null.StringFrom(*request.Description),
		TenantID:    tenantID,
	}
	
	err = claim.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("Failed to create claim")
		return dto.CreateClaimResponse{}, err
	}

	log.Debug().Msg("Claim created successfully")

	return dto.CreateClaimResponse{ID: claim.ID}, nil
}

func (s *Service) Update(ctx context.Context, request dto.UpdateClaimRequest) (dto.UpdateClaimResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "Update").Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.UpdateClaimResponse{}, err
	}

	claim, err := models.Claims(
		models.ClaimWhere.ID.EQ(request.ID),
		models.ClaimWhere.TenantID.EQ(tenantID),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Msg("Claim not found")
			return dto.UpdateClaimResponse{}, httperrors.ErrClaimNotFound
		}

		log.Error().Err(err).Msg("Failed to find claim")
		return dto.UpdateClaimResponse{}, err
	}
	
	changed := false
	if request.Name != nil && claim.Name != *request.Name {
		exists, err := models.Claims(
			models.ClaimWhere.Name.EQ(*request.Name),
			models.ClaimWhere.TenantID.EQ(tenantID),
			models.ClaimWhere.ID.NEQ(request.ID),
		).Exists(ctx, s.db)
		if err != nil {
			log.Error().Err(err).Msg("Failed to check whether claim exists")
			return dto.UpdateClaimResponse{}, err
		}

		if exists {
			log.Debug().Str("name", *request.Name).Msg("Claim already exists")
			return dto.UpdateClaimResponse{}, httperrors.ErrConflictClaimAlreadyExists
		}

		claim.Name = *request.Name
		changed = true
	}

	if request.Description != nil && claim.Description.String != *request.Description {
		claim.Description = null.StringFrom(*request.Description)
		changed = true
	}

	if !changed {
		return dto.UpdateClaimResponse{ID: claim.ID}, nil
	}

	_, err = claim.Update(ctx, s.db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("Failed to update claim")
		return dto.UpdateClaimResponse{}, err
	}

	log.Debug().Msg("Claim updated successfully")

	return dto.UpdateClaimResponse{ID: claim.ID}, nil
}

func (s *Service) Delete(ctx context.Context, request dto.DeleteClaimRequest) (dto.DeleteClaimResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "Delete").Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.DeleteClaimResponse{}, err
	}

	claim, err := models.Claims(
		models.ClaimWhere.ID.EQ(request.ID),
		models.ClaimWhere.TenantID.EQ(tenantID),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error().Err(err).Msg("Claim not found")
			return dto.DeleteClaimResponse{}, httperrors.ErrClaimNotFound
		}

		log.Error().Err(err).Msg("Failed to find claim")
		return dto.DeleteClaimResponse{}, err
	}

	_, err = claim.Delete(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete claim")
		return dto.DeleteClaimResponse{}, err
	}

	log.Debug().Msg("Claim deleted successfully")

	return dto.DeleteClaimResponse{ID: claim.ID}, nil
}
