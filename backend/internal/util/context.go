package util

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"
)

type contextKey string

const (
	CTXKeyUser          contextKey = "user"
	CTXKeyAuthToken     contextKey = "auth_token"
	CTXKeyCacheControl  contextKey = "cache_control"
	CTXKeyRequestID     contextKey = "request_id"
	CTXKeyDisableLogger contextKey = "disable_logger"
	CTXKeyTenant        contextKey = "tenant_id"
)

type detachedContext struct {
	parent context.Context
}

func (c detachedContext) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (c detachedContext) Done() <-chan struct{}             { return nil }
func (c detachedContext) Err() error                        { return nil }
func (c detachedContext) Value(key interface{}) interface{} { return c.parent.Value(key) }

func DetachContext(ctx context.Context) context.Context {
	return detachedContext{ctx}
}

func RequestIDFromContext(ctx context.Context) (string, error) {
	val := ctx.Value(CTXKeyRequestID)
	if val == nil {
		return "", errors.New("no request id present in context")
	}

	id, ok := val.(string)
	if !ok {
		return "", errors.New("request ID in context is not a string")
	}

	return id, nil
}

func ShouldDisableLogger(ctx context.Context) bool {
	s := ctx.Value(CTXKeyDisableLogger)
	if s == nil {
		return false
	}

	shouldDisable, ok := s.(bool)
	if !ok {
		return false
	}

	return shouldDisable
}

func DisableLogger(ctx context.Context, shouldDisable bool) context.Context {
	return context.WithValue(ctx, CTXKeyDisableLogger, shouldDisable)
}

func SaveContextValue(ctx context.Context, key contextKey, value any) context.Context {
	if value == nil {
		return ctx
	}

	valStr := fmt.Sprintf("%v", value)
	return context.WithValue(ctx, key, valStr)
}

func GetContextValue(ctx context.Context, key contextKey) (string, error) {
	val := ctx.Value(key)
	if val == nil {
		return "", errors.New("no value present in context")
	}

	valStr := fmt.Sprintf("%v", val)
	return valStr, nil
}

func TenantIDFromContext(ctx context.Context) (int64, error) {
	valStr, err := GetContextValue(ctx, CTXKeyTenant)
	if err != nil {
		return 0, err
	}

	tenantID, err := strconv.ParseInt(valStr, 10, 64)
	if err != nil {
		return 0, errors.New("tenant id in context is not a valid number")
	}

	return tenantID, nil
}
