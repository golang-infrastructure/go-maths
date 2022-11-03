package maths

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a, b T) T {
	if a >= b {
		return a
	} else {
		return b
	}
}

func Min[T constraints.Ordered](a, b T) T {
	if a <= b {
		return a
	} else {
		return b
	}
}
