// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maths

import "unsafe"

// Float32ToBits returns the IEEE 754 binary representation of f,
// with the sign bit of f and the result in the same bit position.
// Float32ToBits(Float32FromBits(x)) == x.
func Float32ToBits(f float32) uint32 { return *(*uint32)(unsafe.Pointer(&f)) }

// Float32FromBits returns the floating-point number corresponding
// to the IEEE 754 binary representation b, with the sign bit of b
// and the result in the same bit position.
// Float32FromBits(Float32ToBits(x)) == x.
func Float32FromBits(b uint32) float32 { return *(*float32)(unsafe.Pointer(&b)) }

// Float64ToBits returns the IEEE 754 binary representation of f,
// with the sign bit of f and the result in the same bit position,
// and Float64ToBits(Float64FromBits(x)) == x.
func Float64ToBits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }

// Float64FromBits returns the floating-point number corresponding
// to the IEEE 754 binary representation b, with the sign bit of b
// and the result in the same bit position.
// Float64FromBits(Float64ToBits(x)) == x.
func Float64FromBits(b uint64) float64 { return *(*float64)(unsafe.Pointer(&b)) }
