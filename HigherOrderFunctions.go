package main

func Filter[V any](items []V, test func(V) bool) []V {
	results := make([]V, 0)
	for _, item := range items {
		if test(item) {
			results = append(results, item)
		}
	}
	return results
}

func (s *Stream[V]) Filter(test func(V) bool) *Stream[V] {
	stream := NewStream[V]()
	for i := 0; !s.Finished(); i++ {
		if test(s.Get()) {
			stream.Append(s.Get())
		}
		s.Next()
	}
	return &stream
}

func Map[V any, T any](items []V, conversion func(V) T) []T {
	results := make([]T, len(items))
	for i, item := range items {
		results[i] = conversion(item)
	}
	return results
}

func SMap[V any, T any](input Stream[V], conversion func(V) T) *Stream[T] {
	stream := NewStream[T]()
	for !input.Finished() {
		stream.Append(conversion(input.Get()))
		input.Next()
	}
	return &stream
}

func ForEach[V any](items []V, action func(V)) []V {
	for _, item := range items {
		action(item)
	}
	return items
}

func (s *Stream[T]) ForEach(items []T, action func(T)) {
	for !s.Finished() {
		action(s.Get())
		s.Next()
	}
}

func Reduce[V any, A any](items []V, operation func(V, A) A, startingAccumulator A) A {
	for _, item := range items {
		startingAccumulator = operation(item, startingAccumulator)
	}
	return startingAccumulator
}

func SReduce[V any, A any](input Stream[V], operation func(V, A) A, startingAccumulator A) A {
	for !input.Finished() {
		startingAccumulator = operation(input.Get(), startingAccumulator)
		input.Next()
	}
	return startingAccumulator
}

func AnyMatch[V any](items []V, test func(V) bool) bool {
	for _, item := range items {
		if test(item) {
			return true
		}
	}
	return false
}

func (s *Stream[T]) AnyMatch(test func(T) bool) bool {
	for !s.Finished() {
		if test(s.Get()) {
			return true
		}
		s.Next()
	}
	return false
}

func AllMatch[V any](items []V, test func(V) bool) bool {
	for _, item := range items {
		if !test(item) {
			return false
		}
	}
	return true
}

func (s *Stream[T]) AllMatch(test func(T) bool) bool {
	for !s.Finished() {
		if !test(s.Get()) {
			return false
		}
		s.Next()
	}
	return true
}

func NoneMatch[V any](items []V, test func(V) bool) bool {
	for _, item := range items {
		if test(item) {
			return false
		}
	}
	return true
}

func (s *Stream[T]) NoneMatch(test func(T) bool) bool {
	for !s.Finished() {
		if test(s.Get()) {
			return false
		}
		s.Next()
	}
	return true
}

func FindFirst[V any](items []V, test func(V) bool) Optional[V] {
	for _, item := range items {
		if test(item) {
			return Of(item)
		}
	}
	return None[V]()
}

func (s *Stream[T]) FindFirst(test func(T) bool) Optional[T] {
	for !s.Finished() {
		if test(s.Get()) {
			return Of(s.Get())
		}
		s.Next()
	}
	return None[T]()
}
