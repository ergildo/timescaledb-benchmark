package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMedianOddList(t *testing.T) {
	data := []float64{1, 2, 3}
	num := Median(data)
	assert.Equal(t, 2.0, num)
}

func TestMedianEvenList(t *testing.T) {
	data := []float64{1, 2, 3, 4}
	num := Median(data)
	assert.Equal(t, 2.5, num)
}

func TestMedianEmptyList(t *testing.T) {
	var data []float64
	num := Median(data)
	assert.Equal(t, 0.0, num)
}
