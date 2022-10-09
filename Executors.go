package main

type Executor[T any] interface {
	Submit(func() T) Future[T]
}

type PooledExecutor[T any] struct {
	//2 way queues in and out of goroutine
}
