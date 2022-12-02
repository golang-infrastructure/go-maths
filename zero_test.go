package maths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsZero(t *testing.T) {
	t1 := 0.0000000001
	assert.False(t, IsZero(t1))

	t2 := float64(0)
	assert.True(t, IsZero(t2))

	t3 := uint64(0)
	assert.True(t, IsZero(t3))
}
