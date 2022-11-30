package maths

import (
	"unsafe"
)

// Float64SliceToUint64Slice 把64位的浮点型数据转为uint64数据
func Float64SliceToUint64Slice(float64Slice []float64) []uint64 {
	result := make([]uint64, len(float64Slice))
	for index, x := range float64Slice {
		result[index] = *(*uint64)(unsafe.Pointer(&x))
	}
	return result
}

// Uint64SliceToFloat64Slice 把64位的无符号数组转为浮点型数组
func Uint64SliceToFloat64Slice(uint64Slice []uint64) []float64 {
	result := make([]float64, len(uint64Slice))
	for index, x := range uint64Slice {
		result[index] = *(*float64)(unsafe.Pointer(&x))
	}
	return result
}

func Float32SliceToUint32Slice(float32Slice []float32) []uint32 {
	result := make([]uint32, len(float32Slice))
	for index, x := range float32Slice {
		result[index] = *(*uint32)(unsafe.Pointer(&x))
	}
	return result
}

func Uint32SliceToFloat32Slice(uint32Slice []uint32) []float32 {
	result := make([]float32, len(uint32Slice))
	for index, x := range uint32Slice {
		result[index] = *(*float32)(unsafe.Pointer(&x))
	}
	return result
}
