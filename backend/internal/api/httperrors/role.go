package httperrors

import "net/http"

var (
	ErrConflictRoleAlreadyExists = NewHTTPError(http.StatusConflict, "ROLE_ALREADY_EXISTS", "Role with given name already exists")
	ErrRoleNotFound = NewHTTPError(http.StatusNotFound, "ROLE_NOT_FOUND", "Role not found")
)