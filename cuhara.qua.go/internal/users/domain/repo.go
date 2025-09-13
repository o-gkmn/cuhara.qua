package domain

import "context"

type UserRepository interface {
	Create(ctx context.Context, u *User, outEvents []Event) (int64, error)
}