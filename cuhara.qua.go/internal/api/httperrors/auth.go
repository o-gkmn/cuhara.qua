package httperrors

import "net/http"

var (
	ErrConflictUserAlreadyExists = NewHTTPError(http.StatusConflict, "USER_ALREADY_EXISTS", "User with given email already exists")
)
