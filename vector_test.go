package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestNewVector(t *testing.T) {
	vector := NewVector(5, 10)
	assert.Equal(t, vector.Length(), 5)
	assert.Equal(t, vector.Capacity(), 10)
}

func TestVector_Expand(t *testing.T) {
	vector := NewVector(2, 2)
	vector.Expand()
	assert.Equal(t, vector.Capacity(), 4)
}

func TestVector_Shrink(t *testing.T) {
	vector := NewVector(2, 6)
	vector.Shrink()
	assert.Equal(t, vector.Capacity(), 3)
}

func TestVector_Insert(t *testing.T) {
	vector := NewVector(0, 2)
	vector.Insert(0, 1)
	vector.Insert(1, 2)
	vector.Insert(2, 3)
	vector.Insert(3, 4)

	assert.Equal(t, vector.Data(), []int{1, 2, 3, 4})
	assert.Equal(t, vector.Length(), 4)
	assert.Equal(t, vector.Capacity(), 4)

	vector.Insert(1, 6)
	assert.Equal(t, vector.Data(), []int{1, 6, 2, 3, 4})

	vector.Insert(0, 7)
	assert.Equal(t, vector.Data(), []int{7, 1, 6, 2, 3, 4})

	vector.Insert(6, 8)
	assert.Equal(t, vector.Data(), []int{7, 1, 6, 2, 3, 4, 8})
}


func TestVector_DeleteMany(t *testing.T) {
	vector := NewVector(0, 4)
	vector.Insert(0, 1)
	vector.Insert(1, 2)
	vector.Insert(2, 3)
	vector.Insert(3, 4)

	vector.DeleteMany(1, 4)
	assert.Equal(t, vector.Data(), []int{1})

	vector.Insert(1, 2)
	vector.Insert(2, 3)
	vector.Insert(3, 4)

	vector.DeleteMany(1, 3)
	assert.Equal(t, vector.Data(), []int{1, 4})
}

func TestVector_DeleteOne(t *testing.T) {
	vector := NewVector(0, 4)
	vector.Insert(0, 1)
	vector.Insert(1, 2)
	vector.Insert(2, 3)
	vector.Insert(3, 4)

	vector.DeleteOne(2)
	assert.Equal(t, vector.Data(), []int{1, 2, 4})
}
