package domain

import (
	"context"
)

type TennantRepository interface {
	Create(ctx context.Context, t *Tennant, outEvents []Event) (int64, error)
}
