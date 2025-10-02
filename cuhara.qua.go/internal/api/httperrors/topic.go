package httperrors

import "net/http"

var (
	ErrTopicNotFound = NewHTTPError(http.StatusNotFound, "TOPIC_NOT_FOUND", "Topic not found")
	ErrConflictTopicAlreadyExists = NewHTTPError(http.StatusConflict, "TOPIC_ALREADY_EXISTS", "Topic with given name already exists")
	ErrConflictSubTopicAlreadyExists = NewHTTPError(http.StatusConflict, "SUB_TOPIC_ALREADY_EXISTS", "Sub topic with given name already exists")
	ErrSubTopicNotFound = NewHTTPError(http.StatusNotFound, "SUB_TOPIC_NOT_FOUND", "Sub topic not found")
)