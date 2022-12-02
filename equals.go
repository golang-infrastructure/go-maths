package maths

import (
	"github.com/golang-infrastructure/go-gtypes"
	"math"
)

// Equals 判断两个数字是否相等，通过比较距离来判断，当距离足够小的认为是相等，两个参数的类型可以不相等，但只要是数字类型就可以
func Equals[T1, T2 gtypes.Number](a T1, b T2) bool {
	t1 := float64(a)
	t2 := float64(b)
	return math.Abs(t1-t2) < 0.000001
}
