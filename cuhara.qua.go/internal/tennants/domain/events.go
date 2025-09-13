package domain

import "time"

type Event interface {
	EventName() string
	OccurredAt() time.Time
}

type TennantCreated struct {
	ID       int64
	Name     string
	Occurred time.Time
}

func (e *TennantCreated) EventName() string     { return "tennant.TenantCreated" }
func (e *TennantCreated) OccurredAt() time.Time { return e.Occurred }
