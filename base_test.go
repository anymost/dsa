package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestFibonacci(t *testing.T) {
	assert.Equal(t, 55, fibonacci(10))
}

func TestFibonacciBest(t *testing.T) {
	assert.Equal(t, 55, fibonacciBest(10))
}

func TestLCS(t *testing.T) {
	assert.Equal(t, LCSCount([]byte{'a', 'b', 'c', 'd'}, []byte{'a', 'b', 'd', 'e'}), 3)
}

func TestMax2(t *testing.T) {
	assert.Equal(t, max2([]int{7, 6, 5, 4, 3, 2, 1}), []int{7, 6})
}
