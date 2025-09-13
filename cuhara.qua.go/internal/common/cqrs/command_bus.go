package cqrs

import (
	"context"
	"fmt"
)

type CommandBus struct {
	handlers map[string]CommandHandler[any]
}

func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]CommandHandler[any]),
	}
}

func Register[T any](cb *CommandBus, commandType string, handler CommandHandler[T]) {
	cb.handlers[commandType] = &commandHandlerWrapper[T]{handler: handler}
}

func Execute[T any](cb *CommandBus, ctx context.Context, command Command[T]) (T, error) {
	var zero T

	handler, exists := cb.handlers[command.Type()]
	if !exists {
		return zero, fmt.Errorf("handler not found for command: %s", command.Type())
	}

	result, err := handler.Handle(ctx, command)
	if err != nil {
		return zero, err
	}

	return result.(T), nil
}

type commandHandlerWrapper[T any] struct {
	handler CommandHandler[T]
}

func (w *commandHandlerWrapper[T]) Handle(ctx context.Context, command Command[any]) (any, error) {
	return w.handler.Handle(ctx, command.(Command[T]))
}
