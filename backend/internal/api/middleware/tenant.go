package middleware

import (
	"context"
	"database/sql"
	"strconv"

	"cuhara.qua.go/internal/api"
	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/models"
	"cuhara.qua.go/internal/util"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

var (
	skipTenantAuthPaths = []string{"/api/v1/auth/login", "/", "/swagger", "/docs"}
)

const (
	TenantHeader = "X-TENANT-ID"
	MaxTenantID  = 9223372036854775807
	MaxUserID    = 9223372036854775807
	Base10       = 10
	Int64Bits    = 64
)

func TenantAuth(s *api.Server) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log := util.LogFromEchoContext(c).With().Str("middleware", "tenant").Logger()
			ctx := c.Request().Context()

			// skip tenant auth if the request is in the skipTenantAuthPaths
			if skipTenantAuth(c) {
				return next(c)
			}

			// Extract tenant id from headers
			tenantID, err := extractTenantID(c, log)
			if err != nil {
				return err
			}

			// extract user id from context
			userID, err := extractUserID(ctx, log)
			if err != nil {
				return err
			}

			// validate user and tenant related
			err = validateUserTenantRelationship(c, s.DB, userID, tenantID, log)
			if err != nil {
				return err
			}

			ctx = context.WithValue(ctx, util.CTXKeyTenant, tenantID)
			c.SetRequest(c.Request().WithContext(ctx))

			log.Debug().Msg("tenant validation successful")

			return next(c)
		}
	}
}

func skipTenantAuth(c echo.Context) bool {
	for _, path := range skipTenantAuthPaths {
		if c.Request().URL.Path == path {
			return true
		}
	}
	return false
}

func extractTenantID(c echo.Context, log zerolog.Logger) (int64, error) {
	// get tenant id string from header
	tenantIDStr := c.Request().Header.Get(TenantHeader)

	// if tenant id is empty return an error
	if tenantIDStr == "" {
		log.Info().Msg("missing tenant header")
		return 0, httperrors.ErrMissingHeader
	}

	// convert tenant id to int64
	tenantID, err := strconv.ParseInt(tenantIDStr, Base10, Int64Bits)
	if err != nil {
		log.Info().
			Str("tenant_id", tenantIDStr).
			Msg("invalid tenant ID format")
		return 0, httperrors.ErrInvalidHeader
	}

	// check if tenantID overflow max int64 value
	if tenantID > MaxTenantID {
		log.Info().Int64("tenant_id", tenantID).Msg("tenant ID exceeds maximum value")
		return 0, httperrors.ErrInvalidHeader
	}

	// if tenant id lower than zero return an error
	if tenantID <= 0 {
		log.Info().Msg("tenant id equals or lower than zero")
		return 0, httperrors.ErrInvalidHeader
	}

	return tenantID, nil
}

func extractUserID(ctx context.Context, log zerolog.Logger) (int64, error) {
	// extract user ID from context. This user ID was previously added to the context
	// by the JWT middleware that runs before this tenant validation middleware.
	userIDStr, ok := ctx.Value(util.CTXKeyUser).(string)
	if !ok {
		log.Info().Msg("missing user context")
		return 0, httperrors.ErrUnauthorized
	}

	// convert user id to int64
	userID, err := strconv.ParseInt(userIDStr, Base10, Int64Bits)
	if err != nil {
		log.Info().
			Str("user_id", userIDStr).
			Msg("invalid user ID format")
		return 0, httperrors.ErrUnauthorized
	}

	// check if userID overflow max int64 value
	if userID > MaxUserID {
		log.Info().Int64("user_id", userID).Msg("user ID exceeds maximum value")
		return 0, httperrors.ErrInvalidHeader
	}

	// if user id lower than zero return an error
	if userID <= 0 {
		log.Info().Msg("user id equals or lower than zero")
		return 0, httperrors.ErrUnauthorized
	}

	return userID, nil
}

func validateUserTenantRelationship(c echo.Context, db *sql.DB, userID, tenantID int64, log zerolog.Logger) error {
	// find the user with given user id and tenant id
	// this confirms tenant exist and user related this tenant
	exists, err := models.Users(
		models.UserWhere.ID.EQ(userID),
		models.UserWhere.TenantID.EQ(tenantID),
	).Exists(c.Request().Context(), db)
	if err != nil {
		log.Err(err).
			Int64("user_id", userID).
			Int64("tenant_id", tenantID).
			Msg("Database error during tenant validation")

		return err
	}

	if !exists {
		log.Info().
			Int64("user_id", userID).
			Int64("tenant_id", tenantID).
			Msg("user not found with given userID and tenantID")
		return httperrors.ErrForbidden
	}

	return nil
}
