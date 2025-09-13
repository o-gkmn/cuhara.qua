package command

import (
	"context"

	"cuhara.qua.go/internal/common/cqrs"
	"cuhara.qua.go/internal/users/domain"
)

const CreateUserCommandType = "user.create"

type CreateUserResponse struct {
	ID int64 `json:"id" example:"1"`
}

type CreateUserRequest struct {
	Name       string `json:"name" example:"John Doe" binding:"required"`
	Email      string `json:"email" example:"john@example.com" binding:"required,email"`
	VscAccount string `json:"vsc_account" example:"john_doe"`
	RoleID     int64  `json:"role_id" example:"1" binding:"required"`
	TenantID   int64  `json:"tenant_id" example:"1" binding:"required"`
}

type CreateUserCommand struct {
	cqrs.BaseCommand[CreateUserResponse]

	Name       string
	Email      string
	VscAccount string
	RoleID     int64
	TenantID   int64
}

func NewCreateUserCommand(r CreateUserRequest) *CreateUserCommand {
	command := cqrs.NewBaseCommand[CreateUserResponse](CreateUserCommandType)
	return &CreateUserCommand{
		BaseCommand: command,
		Name:        r.Name,
		Email:       r.Email,
		VscAccount:  r.VscAccount,
		RoleID:      r.RoleID,
		TenantID:    r.TenantID,
	}
}

type CreateUserHandler struct {
	repo domain.UserRepository
}

func NewCreateUserHandler(repo domain.UserRepository) *CreateUserHandler {
	return &CreateUserHandler{repo: repo}
}

func (h *CreateUserHandler) Handle(ctx context.Context, c cqrs.Command[CreateUserResponse]) (CreateUserResponse, error) {
	cmd := c.(*CreateUserCommand)

	u, events, err := domain.NewUser(cmd.Name, cmd.Email, cmd.VscAccount, cmd.RoleID, cmd.TenantID)
	if err != nil {
		return CreateUserResponse{}, err
	}

	id, err := h.repo.Create(ctx, u, events)
	if err != nil {
		return CreateUserResponse{}, err
	}

	return CreateUserResponse{ID: id}, nil
}
