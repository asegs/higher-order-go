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

func Map[V any, T any](items []V, conversion func(V) T) []T {
	results := make([]T, len(items))
	for i, item := range items {
		results[i] = conversion(item)
	}
	return results
}

func Reduce[V any, A any](items []V, operation func(V, A) A, startingAccumulator A) A {
	for _, item := range items {
		startingAccumulator = operation(item, startingAccumulator)
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

func AllMatch[V any](items []V, test func(V) bool) bool {
	for _, item := range items {
		if !test(item) {
			return false
		}
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

func FindFirst[V any](items []V, test func(V) bool) Optional[V] {
	for _, item := range items {
		if test(item) {
			return Of(item)
		}
	}
	return None[V]()
}
