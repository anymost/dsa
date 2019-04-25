package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestNewVectorWithData(t *testing.T) {
	vector := NewVectorWithData([]int{1, 2, 3, 4})
	assert.Equal(t, vector.Data(), []int{1, 2, 3, 4})
	assert.Equal(t, vector.Length(), 4)
	assert.Equal(t, vector.Capacity(), 4)
}

func TestNewVector(t *testing.T) {
	vector := NewVector(5, 10)
	assert.Equal(t, vector.Data(), []int{0, 0, 0, 0, 0})
	assert.Equal(t, vector.Length(), 5)
	assert.Equal(t, vector.Capacity(), 10)
}

func TestVector_Expand(t *testing.T) {
	vector := NewVectorWithData([]int{1, 2, 3, 4})
	vector.Expand()
	assert.Equal(t, vector.Capacity(), 8)
}

func TestVector_Shrink(t *testing.T) {
	vector := NewVector(2, 6)
	vector.Shrink()
	assert.Equal(t, vector.Capacity(), 3)
}

func TestVector_Insert(t *testing.T) {
	vector := NewVectorWithData([]int{})
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
	vector := NewVectorWithData([]int{1, 2, 3, 4})

	vector.DeleteMany(1, 4)
	assert.Equal(t, vector.Data(), []int{1})

	vector.Insert(1, 2)
	vector.Insert(2, 3)
	vector.Insert(3, 4)

	vector.DeleteMany(1, 3)
	assert.Equal(t, vector.Data(), []int{1, 4})
}

func TestVector_DeleteOne(t *testing.T) {
	vector := NewVectorWithData([]int{1, 2, 3, 4})
	vector.DeleteOne(2)
	assert.Equal(t, vector.Data(), []int{1, 2, 4})
}

func TestVector_Find(t *testing.T) {
	vector := NewVectorWithData([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	index := vector.Find(5, 0, 10)
	assert.Equal(t, index, 4)
}

func TestVector_BinSearch(t *testing.T) {
	vector := NewVectorWithData([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	index := vector.BinSearch(5, 0, 10)
	assert.Equal(t, index, 4)

	index = vector.BinSearch(10, 0, 10)
	assert.Equal(t, index, 9)

	index = vector.BinSearch(1, 0, 10)
	assert.Equal(t, index, 0)

	index = vector.BinSearch(20, 0, 10)
	assert.Equal(t, index, -1)
}

func TestVector_FibSearch(t *testing.T) {
	vector := NewVectorWithData([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	index := vector.FibSearch(5, 0, 10)
	assert.Equal(t, index, 4)

	index = vector.FibSearch(10, 0, 10)
	assert.Equal(t, index, 9)

	index = vector.FibSearch(1, 0, 10)
	assert.Equal(t, index, 0)
}

func TestVector_BubbleSort(t *testing.T) {
	vector := NewVectorWithData([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	vector.BubbleSort()
	assert.Equal(t, vector.Data(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestVector_MergeSort(t *testing.T) {
	vector := NewVectorWithData([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	vector.MergeSort(0, 10)
	assert.Equal(t, vector.Data(), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
