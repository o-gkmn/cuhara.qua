package command

import (
	"context"

	"cuhara.qua.go/internal/common/cqrs"
	"cuhara.qua.go/internal/tennants/domain"
)

const CreateTennantCommandType = "tennant.create"

type CreateTennantRequest struct {
	Name string `json:"name" example:"KOÇ HOLDİNG AŞ" binding:"required,min=2,max=100"`
}

type CreateTennantResponse struct {
	ID int64 `json:"id"`
}

type CreateTennantCommand struct {
	cqrs.BaseCommand[CreateTennantResponse]

	Name string
}

func NewCreateTennantCommand(r CreateTennantRequest) *CreateTennantCommand {
	command := cqrs.NewBaseCommand[CreateTennantResponse](CreateTennantCommandType)
	return &CreateTennantCommand{
		BaseCommand: command,
		Name:        r.Name,
	}
}

type CreateTennantHandler struct {
	repo domain.TennantRepository
}

func NewCreateTennantHandler(repo domain.TennantRepository) *CreateTennantHandler {
	return &CreateTennantHandler{repo: repo}
}

func (h *CreateTennantHandler) Handle(ctx context.Context, c cqrs.Command[CreateTennantResponse]) (CreateTennantResponse, error) {
	cmd := c.(*CreateTennantCommand)

	t, events, err := domain.NewTennant(cmd.Name)
	if err != nil {
		return CreateTennantResponse{}, err
	}

	id, err := h.repo.Create(ctx, t, events)
	if err != nil {
		return CreateTennantResponse{}, err
	}

	return CreateTennantResponse{ID: id}, nil
}
