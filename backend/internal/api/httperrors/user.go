package httperrors

import "net/http"

var (
	ErrUserNotFound = NewHTTPError(http.StatusNotFound, "USER_NOT_FOUND", "User not found")
	ErrConflictUserEmailAlreadyExists = NewHTTPError(http.StatusConflict, "USER_EMAIL_ALREADY_EXISTS", "User email already exists")
	ErrConflictUserVscAccountAlreadyExists = NewHTTPError(http.StatusConflict, "USER_VSC_ACCOUNT_ALREADY_EXISTS", "User vsc account already exists")
)