package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestNewEmptyList(t *testing.T) {
	list := NewEmptyList()
	assert.Equal(t, list.Head().Next(), list.Tail())
	assert.Equal(t, list.Tail().Prev(), list.Head())
	assert.Equal(t, list.size,0)
}
