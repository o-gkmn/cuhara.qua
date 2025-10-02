package user

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/config"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/models"
	"cuhara.qua.go/internal/util"
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

func (s *Service) GetUsers(ctx context.Context) ([]dto.UserDTO, error) {
	log := util.LogFromContext(ctx).With().Str("function", "GetUsers").Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return nil, err
	}

	users, err := models.Users(
		models.UserWhere.TenantID.EQ(tenantID),
		qm.Load(models.UserRels.Role),
	).All(ctx, s.db)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get users")
		return nil, err
	}

	userDTOs := make([]dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.UserDTO{
			ID:         user.ID,
			Name:       user.Name,
			Email:      user.Email,
			VscAccount: user.VSCAccount,
			RoleDTO: dto.RoleDTO{
				ID:   user.R.Role.ID,
				Name: user.R.Role.Name,
			},
		}
	}

	return userDTOs, nil
}

func (s *Service) Update(ctx context.Context, request dto.UpdateUserRequest) (dto.UpdateUserResponse, error) {
	log := util.LogFromContext(ctx).With().Str("id", strconv.Itoa(int(request.ID))).Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.UpdateUserResponse{}, err
	}

	user, err := models.Users(
		models.UserWhere.ID.EQ(request.ID),
		models.UserWhere.TenantID.EQ(tenantID),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Debug().Err(err).Msg("User not found")
			return dto.UpdateUserResponse{}, httperrors.ErrUserNotFound
		}

		log.Err(err).Msg("Failed to load user")
		return dto.UpdateUserResponse{}, err
	}

	changed := false
	if request.Name != nil && user.Name != *request.Name {
		log.Debug().Str("name", *request.Name).Msg("Updating name")
		user.Name = *request.Name
		changed = true
	}

	if request.Email != nil && user.Email != *request.Email {
		exists, err := models.Users(
			models.UserWhere.Email.EQ(*request.Email),
			models.UserWhere.TenantID.EQ(tenantID),
			models.UserWhere.ID.NEQ(request.ID),
		).Exists(ctx, s.db)
		if err != nil {
			log.Error().Err(err).Msg("Failed to check whether user exists")
			return dto.UpdateUserResponse{}, err
		}

		if exists {
			log.Debug().Str("email", *request.Email).Msg("User email already exists")
			return dto.UpdateUserResponse{}, httperrors.ErrConflictUserEmailAlreadyExists
		}

		log.Debug().Str("email", *request.Email).Msg("Updating email")
		user.Email = *request.Email
		changed = true
	}

	if request.VscAccount != nil && user.VSCAccount != *request.VscAccount {
		exists, err := models.Users(
			models.UserWhere.VSCAccount.EQ(*request.VscAccount),
			models.UserWhere.TenantID.EQ(tenantID),
			models.UserWhere.ID.NEQ(request.ID),
		).Exists(ctx, s.db)
		if err != nil {
			log.Error().Err(err).Msg("Failed to check whether user exists")
			return dto.UpdateUserResponse{}, err
		}

		if exists {
			log.Debug().Str("vsc_account", *request.VscAccount).Msg("User vsc account already exists")
			return dto.UpdateUserResponse{}, httperrors.ErrConflictUserVscAccountAlreadyExists
		}

		log.Debug().Str("vsc_account", *request.VscAccount).Msg("Updating vsc account")
		user.VSCAccount = *request.VscAccount
		changed = true
	}

	if request.RoleID != nil && user.RoleID != *request.RoleID {
		exists, err := models.Roles(
			models.RoleWhere.ID.EQ(*request.RoleID),
			models.RoleWhere.TenantID.EQ(tenantID),
		).Exists(ctx, s.db)
		if err != nil {
			log.Error().Err(err).Msg("Failed to check whether role exists")
			return dto.UpdateUserResponse{}, err
		}

		if !exists {
			log.Debug().Int64("role_id", *request.RoleID).Msg("Role not found")
			return dto.UpdateUserResponse{}, httperrors.ErrRoleNotFound
		}

		log.Debug().Int64("role_id", *request.RoleID).Msg("Updating role id")
		user.RoleID = *request.RoleID
		changed = true
	}

	if !changed {
		return dto.UpdateUserResponse{ID: user.ID}, nil
	}

	_, err = user.Update(ctx, s.db, boil.Whitelist(
		models.UserColumns.Name,
		models.UserColumns.Email,
		models.UserColumns.VSCAccount,
		models.UserColumns.RoleID,
	))
	if err != nil {
		log.Err(err).Msg("Failed to update user")
		return dto.UpdateUserResponse{}, err
	}

	return dto.UpdateUserResponse{ID: user.ID}, nil
}

func (s *Service) Delete(ctx context.Context, request dto.DeleteUserRequest) (dto.DeleteUserResponse, error) {
	log := util.LogFromContext(ctx).With().Int64("id", request.ID).Logger()

	tenantID, err := util.TenantIDFromContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get tenant id from context")
		return dto.DeleteUserResponse{}, err
	}
	
	user, err := models.Users(
		models.UserWhere.ID.EQ(request.ID),
		models.UserWhere.TenantID.EQ(tenantID),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Debug().Err(err).Msg("User not found")
			return dto.DeleteUserResponse{}, httperrors.ErrUserNotFound
		}

		log.Err(err).Msg("Failed to load user")
		return dto.DeleteUserResponse{}, err
	}

	_, err = user.Delete(ctx, s.db)
	if err != nil {
		log.Err(err).Msg("Failed to delete user")
		return dto.DeleteUserResponse{}, err
	}

	return dto.DeleteUserResponse(request), nil
}
