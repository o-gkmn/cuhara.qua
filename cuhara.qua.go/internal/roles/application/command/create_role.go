package command

import (
	"context"

	"cuhara.qua.go/internal/common/cqrs"
	"cuhara.qua.go/internal/roles/domain"
)

const CreateRoleCommandType = "role.create"

type CreateRoleRequest struct {
	Name string `json:"name" example:"ADMIN" binding:"required,min=2,max=100"`
}

type CreateRoleResponse struct {
	ID int64 `json:"id" example:"1"`
}

type CreateRoleCommand struct {
	cqrs.BaseCommand[CreateRoleResponse]

	Name string
}

func NewCreateRoleCommand(r CreateRoleRequest) *CreateRoleCommand {
	command := cqrs.NewBaseCommand[CreateRoleResponse](CreateRoleCommandType)
	return &CreateRoleCommand{
		BaseCommand: command,
		Name:        r.Name,
	}
}

type CreateRoleHandler struct {
	repo domain.RoleRepository
}

func NewCreateRoleHandler(repo domain.RoleRepository) *CreateRoleHandler {
	return &CreateRoleHandler{repo: repo}
}

func (h *CreateRoleHandler) Handle(ctx context.Context, c cqrs.Command[CreateRoleResponse]) (CreateRoleResponse, error) {
	cmd := c.(*CreateRoleCommand)

	r, events, err := domain.NewRole(cmd.Name)
	if err != nil {
		return CreateRoleResponse{}, err
	}

	id, err := h.repo.Create(ctx, r, events)
	if err != nil {
		return CreateRoleResponse{}, err
	}

	return CreateRoleResponse{ID: id}, nil
}
