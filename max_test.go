package maths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	max := Max(1, 2, 3, 4)
	assert.Equal(t, 4, max)
}
