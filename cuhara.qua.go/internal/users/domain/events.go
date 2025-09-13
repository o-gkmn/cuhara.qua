package domain

import "time"

type Event interface {
	EventName() string
	OccurredAt() time.Time
}

type UserCreated struct {
	ID         int64
	Name       string
	Email      string
	VscAccount string
	RoleID     int64
	TenantID   int64
	Occurred   time.Time
}

func (e *UserCreated) EventName() string     { return "users.UserCreated" }
func (e *UserCreated) OccurredAt() time.Time { return e.Occurred }
