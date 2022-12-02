package maths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEquals(t *testing.T) {
	t1 := 0.000001
	t2 := 0.000001
	assert.True(t, Equals(t1, t2))

	t3 := 0.00001
	assert.False(t, Equals(t1, t3))

	t4 := 0
	t5 := float64(0)
	assert.True(t, Equals(t4, t5))
}
