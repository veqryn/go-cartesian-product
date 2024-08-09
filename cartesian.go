package cartesian

import "slices"

// Iter takes interface-slices and returns a channel, receiving cartesian products
func Iter[T any](sets ...[]T) chan []T {
	if len(sets) == 0 {
		c := make(chan []T)
		close(c)
		return c
	}

	buffer := 1
	for _, set := range sets {
		buffer *= len(set)
	}
	if buffer > 1000 {
		buffer = 1000
	}

	c := make(chan []T, buffer)
	go func() {
		defer close(c)
		iterate(c, sets[0], make([]T, 0, len(sets)), sets[1:]...)
	}()
	return c
}

func iterate[T any](c chan []T, topLevel, result []T, needUnpacking ...[]T) {
	if len(needUnpacking) == 0 {
		for _, val := range topLevel {
			c <- append(slices.Clip(result), val)
		}
		return
	}
	for _, val := range topLevel {
		iterate(c, needUnpacking[0], append(result, val), needUnpacking[1:]...)
	}
}
