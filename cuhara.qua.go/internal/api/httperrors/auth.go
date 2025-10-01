package httperrors

import "net/http"

var (
	ErrConflictUserAlreadyExists = NewHTTPError(http.StatusConflict, "USER_ALREADY_EXISTS", "User with given email already exists")
	ErrUnauthorized              = NewHTTPError(http.StatusUnauthorized, "unauthorized", "unauthorized")
	ErrForbidden                 = NewHTTPError(http.StatusForbidden, "forbidden", "forbidden")
	ErrNotFound                  = NewHTTPError(http.StatusNotFound, "not_found", "not_found")
	ErrInvalidToken              = NewHTTPError(http.StatusUnauthorized, "invalid_token", "invalid_token")
	ErrInvalidSubjcet            = NewHTTPError(http.StatusUnauthorized, "invalid_subject", "invalid_subject")
	ErrInvalidIssuer             = NewHTTPError(http.StatusUnauthorized, "invalid_issuer", "invalid_issuer")
	ErrInvalidSigningMethod      = NewHTTPError(http.StatusUnauthorized, "invalid_signing_method", "invalid_signing_method")
)
