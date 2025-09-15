package cqrs

import (
	"context"
	"fmt"
)

type QueryBus struct {
	handlers map[string]QueryHandler[any] 
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[string]QueryHandler[any]),
	}
}

func RegisterQuery[T any](qb *QueryBus, queryType string, handler QueryHandler[T]) {
	qb.handlers[queryType] = &queryHandlerWrapper[T]{handler: handler}
}

func ExecuteQuery[T any](qb *QueryBus, ctx context.Context, query Query[T]) (T, error){
	var zero T

	handler, exists := qb.handlers[query.Type()]
	if !exists {
		return zero, fmt.Errorf("handler not found for query %s", query.Type())
	}

	result, err := handler.Handle(ctx, query)
	if err != nil {
		return zero, err
	}

	return result.(T), nil
}

type queryHandlerWrapper[T any] struct {
	handler QueryHandler[T]
}

func (w *queryHandlerWrapper[T]) Handle(ctx context.Context, query Query[any]) (any, error) {
	return w.handler.Handle(ctx, query.(Query[T]))
}