package domain

import (
	"context"
	"errors"
	"strings"
	"time"
)

var (
	ErrInvalidName = errors.New("isim alanı 2 karakterden büyük 100 karakterden küçük olmalıdır")
)

type Tennant struct {
	ID        int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string     `json:"name" gorm:"not null;uniqueIndex;size:250"`
	CreatedAt time.Time  `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
}

func (Tennant) TableName() string {
	return "tennants"
}

func NewTennant(name string) (*Tennant, []Event, error) {
	t := &Tennant{
		Name:      name,
		CreatedAt: time.Now().UTC(),
	}

	t.Normalize()

	if err := t.Validate(context.Background()); err != nil {
		return nil, nil, err
	}

	ev := &TennantCreated{
		ID:       0,
		Name:     name,
		Occurred: time.Now().UTC(),
	}

	return t, []Event{ev}, nil
}

func (t *Tennant) Normalize() {
	t.Name = strings.TrimSpace(t.Name)
	t.Name = strings.ToUpper(t.Name)
}

func (t *Tennant) Validate(_ context.Context) error {
	if len(t.Name) < 2 || len(t.Name) > 250 {
		return ErrInvalidName
	}
	return nil
}
