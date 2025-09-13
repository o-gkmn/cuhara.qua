package cqrs

import "context"

type Command[T any] interface {
	Type() string
}

type CommandHandler[T any] interface {
	Handle(ctx context.Context, command Command[T]) (T, error)
}

type BaseCommand[T any] struct {
	commandType string
}

func NewBaseCommand[T any](commandType string) BaseCommand[T] {
	return BaseCommand[T]{
		commandType: commandType,
	}
}

func (c BaseCommand[T]) Type() string {
	return c.commandType
}
