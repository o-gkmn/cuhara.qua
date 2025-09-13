package cqrs

import "context"

type Query[T any] interface {
	Type() string
}

type QueryHandler[T any] interface {
	Handle(ctx context.Context, query Query[T]) (T, error)
}

type BaseQuery[T any] struct {
	queryType string
}

func NewBaseQuery[T any](queryType string) BaseQuery[T] {
	return BaseQuery[T]{
		queryType: queryType,
	}
}

func (c BaseQuery[T]) Type() string {
	return c.queryType
}
