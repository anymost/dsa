package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array1 := bubbleSort([]int{5, 4, 3, 2, 1})
	array2 := []int{1, 2, 3, 4, 5}
	assert.Equal(t, array1, array2)
}
