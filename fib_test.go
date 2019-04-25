package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestNewFib(t *testing.T) {
	fib := NewFib(4)
	assert.Equal(t, fib.Get(), 3)
}

func TestFib_Prev(t *testing.T) {
	fib := NewFib(4)
	fib.Prev()
	assert.Equal(t, fib.Get(), NewFib(3).Get())
}
