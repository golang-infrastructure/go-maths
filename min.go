package maths

import "github.com/golang-infrastructure/go-gtypes"

func Min[T gtypes.Number](values ...T) T {
	if len(values) == 0 {
		var zero T
		return zero
	}
	min := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}
