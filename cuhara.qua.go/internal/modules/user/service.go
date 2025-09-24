package user

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"cuhara.qua.go/internal/config"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/models"
	"cuhara.qua.go/internal/util"
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

func (s *Service) GetUsers(ctx context.Context) ([]dto.UserDTO, error) {
	users, err := models.Users().All(ctx, s.db)
	if err != nil {
		return nil, err
	}

	userDTOs := make([]dto.UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.UserDTO{
			ID:         int64(user.ID),
			Name:       user.Name,
			Email:      user.Email,
			VscAccount: user.VSCAccount,
		}
	}

	return userDTOs, nil
}

func (s *Service) Update(ctx context.Context, request dto.UpdateUserRequest) (dto.UpdateUserResponse, error) {
	log := util.LogFromContext(ctx).With().Str("id", strconv.Itoa(int(request.ID))).Logger()

	user, err := models.Users(
		models.UserWhere.ID.EQ(int(request.ID)),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Debug().Err(err).Msg("User not found")
		}

		log.Err(err).Msg("Failed to load user")
		return dto.UpdateUserResponse{}, err
	}

	if request.Name != nil {
		log.Debug().Str("name", *request.Name).Msg("Updating name")
		user.Name = *request.Name
	}

	if request.Email != nil {
		log.Debug().Str("email", *request.Email).Msg("Updating email")
		user.Email = *request.Email
	}

	if request.VscAccount != nil {
		log.Debug().Str("vsc_account", *request.VscAccount).Msg("Updating vsc account")
		user.VSCAccount = *request.VscAccount
	}

	if request.RoleID != nil {
		log.Debug().Int("role_id", int(*request.RoleID)).Msg("Updating role id")
		user.RoleID = int(*request.RoleID)
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

	return dto.UpdateUserResponse{ID: int64(user.ID)}, nil
}

func (s *Service) Delete(ctx context.Context, request dto.DeleteUserRequest) (dto.DeleteUserResponse, error) {
	panic("not implemented")
}
