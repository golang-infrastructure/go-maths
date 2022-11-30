package maths

import "github.com/golang-infrastructure/go-gtypes"

func Max[T gtypes.Number](values ...T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}
	max := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}
