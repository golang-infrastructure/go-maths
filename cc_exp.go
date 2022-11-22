package maths

import (
	"golang.org/x/exp/constraints"
	"math"
	"unsafe"
)

// FloatZero 判断浮点数的值是否为零
func FloatZero[T constraints.Float](floatValue T) bool {
	return *(*uint64)(unsafe.Pointer(&floatValue)) == 0
}

// FloatEquals 比较两个浮点数的值是否相等
func FloatEquals[T constraints.Float](f1, f2 T) bool {
	return math.Abs(float64(f1-f2)) < 0.0000001
}



