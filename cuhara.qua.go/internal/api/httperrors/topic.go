package httperrors

import "net/http"

var (
	ErrTopicNotFound = NewHTTPError(http.StatusNotFound, "TOPIC_NOT_FOUND", "Topic not found")
)