package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"cuhara.qua.go/internal/api/httperrors"
	"cuhara.qua.go/internal/types"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/oapi-codegen/echo-middleware"
	"github.com/rs/zerolog/log"
)

func OpenAPIValidationMiddleware() echo.MiddlewareFunc {
	swagger, err := types.GetSwagger()
	if err != nil {
		log.Error().Err(err).Msg("Failed to get swagger")
		return nil
	}

	opts := &echomiddleware.Options{
		SilenceServersWarning: true,
		Skipper:               skipValidation,
		ErrorHandler: func(c echo.Context, err *echo.HTTPError) error {
			return convertValidationError(err)
		},
		Options: openapi3filter.Options{
			AuthenticationFunc: authenticationFunc,
		},
	}

	return echomiddleware.OapiRequestValidatorWithOptions(swagger, opts)
}

func authenticationFunc(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
	return nil
}

func skipValidation(c echo.Context) bool {
	path := c.Request().URL.Path
	return path == "/swagger" || path == "/docs" || strings.HasPrefix(path, "/swagger/")
}

func convertValidationError(err error) error {
	var reqErr *openapi3filter.RequestError
	log.Error().Err(err).Msg("Request error type: " + reflect.TypeOf(err).String())
	if !errors.As(err, &reqErr) {
		return httperrors.NewHTTPError(
			http.StatusBadRequest,
			httperrors.HTTPErrorTypeGeneric,
			"Invalid request",
		)
	}

	if reqErr.Parameter != nil {
		return handleParameterError(reqErr)
	}

	if reqErr.RequestBody != nil {
		return handleBodyError(reqErr)
	}

	return httperrors.NewHTTPError(
		http.StatusBadRequest,
		httperrors.HTTPErrorTypeGeneric,
		reqErr.Error(),
	)
}

func handleParameterError(reqErr *openapi3filter.RequestError) error {
	fieldName := reqErr.Parameter.Name
	location := reqErr.Parameter.In

	log.Error().Any("reqErr", reqErr).Msg("handleParameterError")

	return httperrors.NewHTTPValidationError(
		http.StatusBadRequest,
		httperrors.HTTPErrorTypeGeneric,
		http.StatusText(http.StatusBadRequest),
		[]types.HttpValidationErrorDetail{
			{Key: fieldName, In: location, Error: fmt.Sprintf("%s parametresi geçersiz.", reqErr.Parameter.Name)},
		},
	)
}

func handleBodyError(reqErr *openapi3filter.RequestError) error {
	var se *openapi3.SchemaError
	if !errors.As(reqErr.Err, &se) {
		return httperrors.NewHTTPValidationError(
			http.StatusBadRequest,
			httperrors.HTTPErrorTypeGeneric,
			fmt.Sprintf("Body validation failed: %s", reqErr.Error()),
			[]types.HttpValidationErrorDetail{},
		)
	}

	jsonPointer := se.JSONPointer()
	fieldName := "body"
	if len(jsonPointer) > 0 {
		fieldName = jsonPointer[len(jsonPointer)-1]
	}

	customMsgs := getCustomErrorMessagesFromProperty(se)
	errorMsg := se.Reason
	if len(customMsgs) > 0 {
		errorMsg = strings.Join(customMsgs, "; ")
	}

	return httperrors.NewHTTPValidationError(
		http.StatusBadRequest,
		httperrors.HTTPErrorTypeGeneric,
		"Body validation failed",
		[]types.HttpValidationErrorDetail{
			{
				Key:   fieldName,
				In:    "body",
				Error: errorMsg,
			},
		},
	)
}

func getCustomErrorMessagesFromProperty(se *openapi3.SchemaError) []string {
	msgs := []string{}
	
	if se.Schema == nil {
		return []string{se.Reason}
	}

	parentSchema := se.Schema
	jsonPointer := se.JSONPointer()
	
	var propertySchema *openapi3.Schema
	
	if len(jsonPointer) > 0 {
		propertyName := jsonPointer[len(jsonPointer)-1]
		
		if parentSchema.Properties != nil {
			if propSchemaRef, exists := parentSchema.Properties[propertyName]; exists {
				if propSchemaRef.Value != nil {
					propertySchema = propSchemaRef.Value
				}
			}
		}
	}
	
	if propertySchema == nil {
		propertySchema = parentSchema
	}
	
	ext, ok := propertySchema.Extensions["x-error-messages"]
	if !ok {
		return msgs
	}


	errorMessages, ok := ext.(map[string]interface{})
	if !ok {
		return msgs
	}

	validationType := detectValidationType(se, propertySchema)
	
	if customMsg, exists := errorMessages[validationType]; exists {
		if msg, ok := customMsg.(string); ok {
			msgs = append(msgs, msg)
		}
	}

	if len(msgs) == 0 {
		for _, key := range []string{"default", "generic", "error"} {
			if customMsg, exists := errorMessages[key]; exists {
				if msg, ok := customMsg.(string); ok {
					msgs = append(msgs, msg)
					break
				}
			}
		}
	}

	return msgs
}

func detectValidationType(se *openapi3.SchemaError, s *openapi3.Schema) string {
	reason := strings.ToLower(se.Reason)

	if strings.Contains(reason, "missing") || strings.Contains(reason, "required") {
		return "required"
	}

	switch {
	case strings.Contains(reason, "minimum") || s.Min != nil:
		return "minimum"
	case strings.Contains(reason, "maximum") || s.Max != nil:
		return "maximum"
	case strings.Contains(reason, "minlength") || s.MinLength > 0:
		return "minLength"
	case strings.Contains(reason, "maxlength") || s.MaxLength != nil:
		return "maxLength"
	case strings.Contains(reason, "pattern") || s.Pattern != "":
		return "pattern"
	case strings.Contains(reason, "format"):
		return "format"
	case strings.Contains(reason, "type"):
		return "type"
	case strings.Contains(reason, "enum"):
		return "enum"
	default:
		return "default"
	}
}