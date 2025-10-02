package httperrors

import "net/http"

var (
	ErrClaimNotFound = NewHTTPError(http.StatusNotFound, "CLAIM_NOT_FOUND", "Claim not found")
	ErrConflictClaimAlreadyExists = NewHTTPError(http.StatusConflict, "CLAIM_ALREADY_EXISTS", "Claim with given name already exists")
)