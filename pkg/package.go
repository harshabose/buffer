package buffer

import "context"

type Pool[T any] interface {
	Get() *T
	Put(*T)
	Release()
}

type Buffer[T any] interface {
	Push(context.Context, *T) error
	Pop(ctx context.Context) (*T, error)
	Size() int
}

type BufferWithGenerator[T any] interface {
	Push(context.Context, *T) error
	Pop(ctx context.Context) (*T, error)
	Size() int
	Generate() *T
	PutBack(*T)
	GetChannel() chan *T
}
