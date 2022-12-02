package maths

import "github.com/golang-infrastructure/go-gtypes"

func Avg[T gtypes.Number](slice []T) float64 {
	return AvgR[T, float64](slice)
}

func AvgR[T, R gtypes.Number](slice []T) R {
	if len(slice) == 0 {
		return 0
	}
	sum := float64(0)
	for _, value := range slice {
		sum += float64(value)
	}
	return R(sum / float64(len(slice)))
}
