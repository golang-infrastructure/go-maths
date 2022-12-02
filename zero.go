package maths

import (
	"github.com/golang-infrastructure/go-gtypes"
	"math"
)

// IsZero 判断是否为0
func IsZero[T gtypes.Number](value T) bool {
	return math.Float64bits(float64(value)) == 0
}
