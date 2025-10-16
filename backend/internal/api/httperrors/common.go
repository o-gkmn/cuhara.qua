package httperrors

import "net/http"

var (
	ErrInvalidID      = NewHTTPError(http.StatusBadRequest, "INVALID_ID", "Invalid ID")
	ErrMissingHeader  = NewHTTPError(http.StatusUnauthorized, "MISSING_HEADER", "missing_header")
	ErrInvalidHeader  = NewHTTPError(http.StatusUnauthorized, "INVALID_HEADER", "invalid_header")
	ErrInternalServer = NewHTTPError(http.StatusInternalServerError, "INTERNAL_SERVER", "Internal Server Error")
)
