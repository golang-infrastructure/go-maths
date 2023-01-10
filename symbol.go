package maths

import "github.com/golang-infrastructure/go-gtypes"

// OppositeNumber 取相反数
func OppositeNumber[T gtypes.Signed](v T) T {
	return -v
}
