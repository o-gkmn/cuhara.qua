package domain

import "time"

type Event interface {
	EventName() string
	OccurredAt() time.Time
}

type RoleCreated struct {
	ID       int64
	Name     string
	Occurred time.Time
}

func (e *RoleCreated) EventName() string { return "roles.RolesCreated" }
func (e *RoleCreated) OccurredAt() time.Time { return e.Occurred }
