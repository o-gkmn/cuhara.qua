package auth

import (
	"context"
	"database/sql"
	"errors"

	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/config"
	"cuhara.qua.go/internal/data/dto"
	"cuhara.qua.go/internal/models"
	"cuhara.qua.go/internal/util"
	"cuhara.qua.go/internal/util/db"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Service struct {
	config config.Server
	db     *sql.DB
}

func NewService(config config.Server, db *sql.DB) *Service {
	return &Service{
		config: config,
		db:     db,
	}
}

// Login implements infra.AuthService.
func (s *Service) Login(ctx context.Context, request dto.LoginRequest) (dto.LoginResponse, error) {
	log := util.LogFromContext(ctx).With().Str("email", request.Email).Logger()

	user, err := models.Users(
		models.UserWhere.Email.EQ(request.Email),
	).One(ctx, s.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Debug().Err(err).Msg("User not found")
		}

		log.Err(err).Msg("Failed to load user")
	}

	matches, err := util.ComparePasswordAndHash(request.Password, user.Password)
	if err != nil {
		log.Err(err).Msg("Failed to compare password with stored hash")
		return dto.LoginResponse{}, err
	}

	if !matches {
		log.Debug().Msg("Provided password does not match stored hash")
		return dto.LoginResponse{}, echo.ErrUnauthorized
	}

	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
	}

	token, err := util.GenerateJWT(claims)
	if err != nil {
		log.Err(err).Msg("Failed to generate JWT")
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{Token: token}, nil
}

// Register implements infra.AuthService.
func (s *Service) Register(ctx context.Context, request dto.RegisterRequest) (dto.LoginResponse, error) {
	log := util.LogFromContext(ctx).With().Str("email", request.Email).Logger()

	exists, err := models.Users(
		models.UserWhere.Email.EQ(request.Email),
	).Exists(ctx, s.db)
	if err != nil {
		log.Err(err).Msg("Failed to check whether user exists")
		return dto.LoginResponse{}, err
	}

	if exists {
		log.Debug().Msg("User with given email already exists")
		return dto.LoginResponse{}, httperrors.ErrConflictUserAlreadyExists
	}

	hash, err := util.HashPassword(request.Password, util.DefaultArgon2Params)
	if err != nil {
		log.Err(err).Msg("Failed to hash user password")
		return dto.LoginResponse{}, err
	}

	var result dto.LoginResponse
	db.WithTransaction(ctx, s.db, func(ce boil.ContextExecutor) error {
		user := &models.User{
			Email:      request.Email,
			Name:       request.Name,
			VSCAccount: request.VscAccount,
			Password:   hash,
			RoleID:     request.RoleID,
			TenantID:   request.TenantID,
		}

		if err := user.Insert(ctx, ce, boil.Infer()); err != nil {
			log.Err(err).Msg("Failed to insert user")
			return err
		}

		claims := jwt.MapClaims{
			"sub":   user.ID,
			"email": user.Email,
		}

		result.Token, err = util.GenerateJWT(claims)
		if err != nil {
			log.Err(err).Msg("Failed to generate JWT")
			return err
		}

		return nil
	})

	return result, nil
}
