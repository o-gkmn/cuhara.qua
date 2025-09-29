package httperrors

import "net/http"

var (
	ErrClaimNotFound = NewHTTPError(http.StatusNotFound, "CLAIM_NOT_FOUND", "Claim not found")
)