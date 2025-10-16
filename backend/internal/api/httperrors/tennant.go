package httperrors

import "net/http"

var (
	ErrConflitTenantAlreadyExist = NewHTTPError(http.StatusConflict, "TENANT_ALREADY_EXIST", "Tenant with given name already exist")
	ErrTenantNotFound = NewHTTPError(http.StatusNotFound, "TENANT_NOT_FOUND", "Tenant not found")
)
