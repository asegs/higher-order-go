package main

import "time"

type Future[T any] interface {
	Done() bool
	Get() T
	GetTimeout(t time.Duration) Optional[T]
	Complete(v T)
}

type AsyncFuture[T any] struct {
	channel  chan bool
	Value    T
	Finished bool
}

type CompletableFuture[T any] struct {
	channel  chan bool
	Value    T
	Finished bool
}

func (f *AsyncFuture[T]) Complete(v T) {
	f.channel <- true
	f.Value = v
	f.Finished = true
}

func ComputeAsync[T any](f func() T) *AsyncFuture[T] {
	channel := make(chan bool, 1)
	future := &AsyncFuture[T]{channel: channel, Finished: false}
	go future.Complete(f())
	return future
}

func (f *AsyncFuture[T]) Done() bool {
	return f.Finished
}

func (f *AsyncFuture[T]) Get() T {
	if f.Done() {
		return f.Value
	}
	<-f.channel
	return f.Value
}

func (f *AsyncFuture[T]) GetWithTimeout(duration time.Duration) Optional[T] {
	if f.Done() {
		return Of(f.Value)
	}
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(duration)
		timeout <- true
	}()
	select {
	case <-timeout:
		return None[T]()
	case <-f.channel:
		return Of(f.Value)
	}
}

func (f *CompletableFuture[T]) Complete(v T) {
	f.channel <- true
	f.Value = v
	f.Finished = true
}

func ComputeCompletable[T any](f func() T) *CompletableFuture[T] {
	channel := make(chan bool, 1)
	return &CompletableFuture[T]{channel: channel}
}

func (f *CompletableFuture[T]) Done() bool {
	return f.Finished
}

func (f *CompletableFuture[T]) Get() T {
	if f.Done() {
		return f.Value
	}
	<-f.channel
	return f.Value
}

func (f *CompletableFuture[T]) GetWithTimeout(duration time.Duration) Optional[T] {
	if f.Done() {
		return Of(f.Value)
	}
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(duration)
		timeout <- true
	}()
	select {
	case <-timeout:
		return None[T]()
	case <-f.channel:
		return Of(f.Value)
	}
}
