package maths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMin(t *testing.T) {
	min := Min(1, 2, 3, 4)
	assert.Equal(t, 1, min)
}
