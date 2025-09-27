package httperrors

import "net/http"

var (
	ErrInvalidID = NewHTTPError(http.StatusBadRequest, "INVALID_ID", "Invalid ID")
)