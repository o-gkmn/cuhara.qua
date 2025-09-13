package domain

import (
	"context"
	"errors"
	"strings"
	"time"

	roledomain "cuhara.qua.go/internal/roles/domain"
	tennantdomain "cuhara.qua.go/internal/tennants/domain"
)

var (
	ErrInvalidName  = errors.New("isim alanı 2 karakterden büyük 100 karakterden küçük olmalıdır")
	ErrInvalidEmail = errors.New("email formatı hatalı")
)

type User struct {
	ID         int64                  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string                 `json:"name" gorm:"size:100;not null"`
	Email      string                 `json:"email" gorm:"size:255;uniqueIndex;not null"`
	VscAccount string                 `json:"vscAccount" gorm:"column:vsc_account;size:255"`
	RoleID     int64                  `json:"roleId" gorm:"column:role_id;not null"`
	TenantID   int64                  `json:"tenantId" gorm:"column:tenant_id;not null"`
	CreatedAt  time.Time              `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  *time.Time             `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
	Role       *roledomain.Role       `json:"role" gorm:"foreignKey:RoleID;references:ID"`
	Tennant    *tennantdomain.Tennant `json:"tennant" gorm:"foreignKey:TenantID;references:ID"`
}

func (User) TableName() string {
	return "users"
}

func NewUser(name, email, vscAccount string, roleID, tenantID int64) (*User, []Event, error) {
	u := &User{
		Name:       name,
		Email:      email,
		VscAccount: vscAccount,
		RoleID:     roleID,
		TenantID:   tenantID,
		CreatedAt:  time.Now().UTC(),
	}

	u.Normalize()

	if err := u.Validate(context.Background()); err != nil {
		return nil, nil, err
	}

	ev := &UserCreated{
		ID:         0,
		Name:       u.Name,
		Email:      u.Email,
		VscAccount: u.VscAccount,
		RoleID:     u.RoleID,
		TenantID:   u.TenantID,
		Occurred:   time.Now().UTC(),
	}

	return u, []Event{ev}, nil
}

func (u *User) Validate(_ context.Context) error {
	if len(u.Name) < 2 || len(u.Name) > 100 {
		return ErrInvalidName
	}
	if !strings.Contains(u.Email, "@") {
		return ErrInvalidEmail
	}

	return nil
}

func (u *User) Normalize() {
	u.Name = strings.TrimSpace(u.Name)
	u.Name = strings.ToUpper(u.Name)

	u.Email = strings.TrimSpace(u.Email)
	u.Email = strings.ToLower(u.Email)

	u.VscAccount = strings.TrimSpace(u.VscAccount)
	u.VscAccount = strings.ToLower(u.VscAccount)
}
