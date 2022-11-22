// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maths

// Signbit reports whether x is negative or negative zero.
func Signbit(x float64) bool {
	return Float64ToBits(x)&(1<<63) != 0
}
