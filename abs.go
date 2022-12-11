package maths

import (
	"github.com/golang-infrastructure/go-gtypes"
	"math"
)

func Abs[T gtypes.Number](v T) T {
	return T(math.Abs(float64(v)))
}
