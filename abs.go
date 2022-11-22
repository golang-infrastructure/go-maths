// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maths

import (
	"golang.org/x/exp/constraints"
)

// Abs returns the absolute value of x.
//
// Special cases are:
//
//	Abs(±Inf) = +Inf
//	Abs(NaN) = NaN
func Abs[T constraints.Float | constraints.Signed](v T) T {
	return T(Float64FromBits(Float64ToBits(float64(v)) &^ (1 << 63)))
}
