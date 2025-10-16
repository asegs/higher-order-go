package main

import "fmt"

type StreamNode[T any] struct {
	Value T
	Next  *StreamNode[T]
}

type Stream[T any] struct {
	Head    *StreamNode[T]
	Tail    *StreamNode[T]
	Current *StreamNode[T]
	Length  int
}

func (s *Stream[T]) Get() T {
	return s.Current.Value
}

func (s *Stream[T]) HasNext() bool {
	return s.Current != nil && s.Current.Next != nil
}

func (s *Stream[T]) Next() {
	if s.Current != nil {
		s.Current = s.Current.Next
	}
}

func (s *Stream[T]) Len() int {
	return s.Length
}

func (s *Stream[T]) Append(item T) {
	newTail := &StreamNode[T]{Value: item, Next: nil}
	if s.Head == nil {
		s.Head = newTail
		s.Current = newTail
		s.Tail = newTail
	} else {
		s.Tail.Next = newTail
		s.Tail = newTail
	}
	s.Length++
}

func (s *Stream[T]) Reset() {
	s.Current = s.Head
}

func (s *Stream[T]) Terminate() {
	s.Current = s.Tail
}

func (s *Stream[T]) ToSlice() []T {
	result := make([]T, s.Len())
	for i := 0; !s.Finished(); i++ {
		result[i] = s.Get()
		s.Next()
	}
	return result
}

func (s *Stream[T]) Finished() bool {
	return s.Len() == 0 || s.Current == nil
}

func NewStream[T any]() Stream[T] {
	return Stream[T]{Head: nil, Current: nil, Tail: nil, Length: 0}
}

func StreamFrom[T any](items []T) Stream[T] {
	stream := NewStream[T]()
	if len(items) == 0 {
		return stream
	}
	for _, item := range items {
		stream.Append(item)
	}
	return stream
}

func SplitStream[T any](stream Stream[T], into int) []Stream[T] {
	streams := make([]Stream[T], into)
	for i := 0; i < into; i++ {
		streams[i] = NewStream[T]()
	}
	for i := 0; !stream.Finished(); i++ {
		fmt.Println(stream.Get())
		streams[i%into].Append(stream.Get())
		stream.Next()
	}
	return streams
}
