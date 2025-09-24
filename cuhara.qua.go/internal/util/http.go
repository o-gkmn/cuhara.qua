package util

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/types"
	oerrors "github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/labstack/echo/v4"
)

func BindAndValidateBody(c echo.Context, v runtime.Validatable) error {
	binder := c.Echo().Binder.(*echo.DefaultBinder)

	if err := binder.BindBody(c, v); err != nil {
		return err
	}

	return validatePayload(c, v)
}

func BindAndValidatePathAndQueryParams(c echo.Context, v runtime.Validatable) error {
	binder := c.Echo().Binder.(*echo.DefaultBinder)

	if err := binder.BindPathParams(c, v); err != nil {
		return err
	}

	if err := binder.BindQueryParams(c, v); err != nil {
		return err
	}

	return validatePayload(c, v)
}

func BindValidatePathParams(c echo.Context, v runtime.Validatable) error {
	binder := c.Echo().Binder.(*echo.DefaultBinder)

	if err := binder.BindPathParams(c, v); err != nil {
		return err
	}

	return validatePayload(c, v)
}

func BindValidateQueryParams(c echo.Context, v runtime.Validatable) error {
	binder := c.Echo().Binder.(*echo.DefaultBinder)

	if err := binder.BindQueryParams(c, v); err != nil {
		return err
	}

	return validatePayload(c, v)
}

func ValidateAndReturn(c echo.Context, code int, v runtime.Validatable) error {
	if err := validatePayload(c, v); err != nil {
		return err
	}

	return c.JSON(code, v)
}

func validatePayload(c echo.Context, v runtime.Validatable) error {
	if err := v.Validate(strfmt.Default); err != nil {

		var compositeError *oerrors.CompositeError
		if errors.As(err, &compositeError) {
			LogFromEchoContext(c).Debug().Errs("validation_errors", compositeError.Errors).Msg("Payload did match scheme, returning HTTP validation error")

			valErrs := formatValidationErrors(c.Request().Context(), compositeError)

			return httperrors.NewHTTPValidationError(http.StatusBadRequest, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusBadRequest), valErrs)
		}

		var validationError *oerrors.Validation
		if errors.As(err, &validationError) {
			LogFromEchoContext(c).Debug().AnErr("validation_error", validationError).Msg("Payload did match scheme, returning HTTP validation error")

			valErrs := []*types.HTTPValidationErrorDetail{
				{
					Key:   &validationError.Name,
					In:    &validationError.In,
					Error: swag.String(validationError.Error()),
				},
			}

			return httperrors.NewHTTPValidationError(http.StatusBadRequest, httperrors.HTTPErrorTypeGeneric, http.StatusText(http.StatusBadRequest), valErrs)
		}

		LogFromEchoContext(c).Error().Err(err).Msg("Failed to validate payload, returning generic HTTP error")
		return err
	}

	return nil
}

func formatValidationErrors(ctx context.Context, err *oerrors.CompositeError) []*types.HTTPValidationErrorDetail {
	valErrs := make([]*types.HTTPValidationErrorDetail, 0, len(err.Errors))
	for _, e := range err.Errors {
		var compositeError *oerrors.CompositeError
		if errors.As(e, &compositeError) {
			valErrs = append(valErrs, formatValidationErrors(ctx, compositeError)...)
			continue
		}

		var validationError *oerrors.Validation
		if errors.As(e, &validationError) {
			valErrs = append(valErrs, &types.HTTPValidationErrorDetail{
				Key:   &validationError.Name,
				In:    &validationError.In,
				Error: swag.String(validationError.Error()),
			})
		}

		LogFromContext(ctx).Warn().Err(e).Str("err_type", fmt.Sprintf("%T", e)).Msg("Received unknown error type while validating payload, skipping")
	}

	return valErrs
}
