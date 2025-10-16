package util

import (
	"context"
	"errors"
	"net/http"

	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/types"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func BindAndValidateBody(c echo.Context, v interface{}) error {
	binder := c.Echo().Binder.(*echo.DefaultBinder)

	if err := binder.BindBody(c, v); err != nil {
		return err
	}

	return validatePayload(c, v)
}

func BindAndValidatePathAndQueryParams(c echo.Context, v interface{}) error {
	binder := c.Echo().Binder.(*echo.DefaultBinder)

	if err := binder.BindPathParams(c, v); err != nil {
		return err
	}

	if err := binder.BindQueryParams(c, v); err != nil {
		return err
	}

	return validatePayload(c, v)
}

func BindValidatePathParams(c echo.Context, v interface{}) error {
	binder := c.Echo().Binder.(*echo.DefaultBinder)

	if err := binder.BindPathParams(c, v); err != nil {
		return err
	}

	return validatePayload(c, v)
}

func BindValidateQueryParams(c echo.Context, v interface{}) error {
	binder := c.Echo().Binder.(*echo.DefaultBinder)

	if err := binder.BindQueryParams(c, v); err != nil {
		return err
	}

	return validatePayload(c, v)
}

func ValidateAndReturn(c echo.Context, code int, v interface{}) error {
	if err := validatePayload(c, v); err != nil {
		return err
	}

	return c.JSON(code, v)
}

func validatePayload(c echo.Context, v interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.StructCtx(c.Request().Context(), v); err != nil {
		valErrs := formatValidationErrors(c.Request().Context(), err)

		return httperrors.NewHTTPValidationError(
			http.StatusBadRequest,
			httperrors.HTTPErrorTypeGeneric,
			http.StatusText(http.StatusBadRequest),
			valErrs,
		)

	}

	return nil
}

func formatValidationErrors(ctx context.Context, err error) []types.HttpValidationErrorDetail {
	var verrs validator.ValidationErrors
	if errors.As(err, &verrs) {
		LogFromContext(ctx).Debug().Err(err).Msg("Payload validation failed, returning HTTP validation error")

		valErrs := make([]types.HttpValidationErrorDetail, 0, len(verrs))
		for _, fe := range verrs {
			valErrs = append(valErrs, types.HttpValidationErrorDetail{
				Key:   fe.Field(),
				In:    "body",
				Error: fe.Error(),
			})
		}

		return valErrs
	}

	return nil
}
