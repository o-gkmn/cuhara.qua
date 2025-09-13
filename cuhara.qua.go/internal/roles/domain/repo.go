package domain

import "context"

type RoleRepository interface {
	Create(ctx context.Context, r *Role, outEvents []Event) (int64, error)
}
