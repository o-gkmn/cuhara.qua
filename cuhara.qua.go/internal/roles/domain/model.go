package domain

import (
	"context"
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidName = errors.New("isim alanı 2 karakterden küçük yada 100 karakterden büyük olamaz")
)

type Role struct {
	ID        int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" gorm:"uniqueIndex;not null;size:255"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
}

func (Role) TableName() string {
	return "roles"
}

func NewRole(name string) (*Role, []Event, error) {
	r := &Role{
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}

	r.Normalize()

	if err := r.Validate(context.Background()); err != nil {
		return nil, nil, err
	}

	ev := &RoleCreated{
		ID:       0,
		Name:     r.Name,
		Occurred: time.Now().UTC(),
	}

	return r, []Event{ev}, nil
}

func (r *Role) Validate(_ context.Context) error {
	if len(r.Name) < 2 || len(r.Name) > 100 {
		return ErrInvalidName
	}
	return nil
}

func (r *Role) Normalize() {
	r.Name = strings.TrimSpace(r.Name)
	r.Name = strings.ToUpper(r.Name)
}
