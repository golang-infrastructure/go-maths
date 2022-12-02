package maths

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAvg(t *testing.T) {

	slice := []float64{
		0.1, 0.2, 0.3, 0.4,
	}
	r := AvgR[float64, float64](slice)
	assert.Equal(t, 0.25, r)

}
