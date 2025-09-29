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

	claims, err := models.Claims().All(ctx, s.db)
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

	return claimDTOs, nil
}

func (s *Service) Create(ctx context.Context, request dto.CreateClaimRequest) (dto.CreateClaimResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "Create").Logger()

	claim := models.Claim{
		Name:        request.Name,
		Description: null.StringFrom(*request.Description),
	}
	
	err := claim.Insert(ctx, s.db, boil.Infer())
	if err != nil {
		log.Error().Err(err).Msg("Failed to create claim")
		return dto.CreateClaimResponse{}, err
	}

	return dto.CreateClaimResponse{ID: claim.ID}, nil
}

func (s *Service) Update(ctx context.Context, request dto.UpdateClaimRequest) (dto.UpdateClaimResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "Update").Logger()

	claim, err := models.FindClaim(ctx, s.db, request.ID)
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

	return dto.UpdateClaimResponse{ID: claim.ID}, nil
}

func (s *Service) Delete(ctx context.Context, request dto.DeleteClaimRequest) (dto.DeleteClaimResponse, error) {
	log := util.LogFromContext(ctx).With().Str("function", "Delete").Logger()

	claim, err := models.FindClaim(ctx, s.db, request.ID)
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

	return dto.DeleteClaimResponse{ID: claim.ID}, nil
}
