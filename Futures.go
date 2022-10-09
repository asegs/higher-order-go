package main

import "time"

type Future[T any] struct {
	channel  chan bool
	Value    T
	Finished bool
}

func Compute[T any](f func() T) *Future[T] {
	channel := make(chan bool, 1)
	future := &Future[T]{channel: channel, Finished: false}
	go func() {
		v := f()
		future.channel <- true
		future.Value = v
		future.Finished = true
	}()
	return future
}

func (f *Future[T]) Done() bool {
	return f.Finished
}

func (f *Future[T]) Get() T {
	if f.Done() {
		return f.Value
	}
	<-f.channel
	return f.Value
}

func (f *Future[T]) GetWithTimeout(duration time.Duration) Optional[T] {
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
