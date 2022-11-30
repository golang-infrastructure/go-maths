package maths

import (
	"testing"
)

func TestFloat64SliceToUint64Slice(t *testing.T) {
	slice := []float64{
		1.0,
		3.14,
	}
	r := Float64SliceToUint64Slice(slice)
	t.Log(r)
}

func TestUint64SliceToFloat64Slice(t *testing.T) {
	slice := []uint64{
		1,
		2,
	}
	r := Uint64SliceToFloat64Slice(slice)
	t.Log(r)
}
