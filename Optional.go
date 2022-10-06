package main

type Optional[T any] interface {
	Get() T
	IsFull() bool
	IsEmpty() bool
	GetOr(d T) T
}

type Full[T any] struct {
	v T
}

func (s Full[T]) Get() T {
	return s.v
}

func (s Full[T]) IsFull() bool {
	return true
}

func (s Full[T]) IsEmpty() bool {
	return false
}

func (s Full[T]) GetOr(d T) T {
	return s.v
}

type Empty[T any] struct{}

func (s Empty[T]) Get() T {
	panic("tried to get empty")
}

func (s Empty[T]) IsFull() bool {
	return false
}

func (s Empty[T]) IsEmpty() bool {
	return true
}

func (s Empty[T]) GetOr(d T) T {
	return d
}

func Of[T any](v T) Optional[T] {
	return Full[T]{v}
}

func None[T any]() Optional[T] {
	return Empty[T]{}
}
