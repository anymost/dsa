package main

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestNewEmptyList(t *testing.T) {
	list := NewEmptyList()
	assert.Equal(t, list.Head().Next(), list.Tail())
	assert.Equal(t, list.Tail().Prev(), list.Head())
	assert.Equal(t, list.size, 0)
}

func TestNewListWithArray(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	list := NewListWithArray(array)
	head := list.Head()
	for i := 0; i < list.size; i++ {
		head = head.Next()
		assert.Equal(t, head.Data(), i+1)
	}
}

func TestList_Head(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	list := NewListWithArray(array)
	assert.Equal(t, list.Head().Next().Data(), 1)
}

func TestList_Tail(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	list := NewListWithArray(array)
	assert.Equal(t, list.Tail().prev.Data(), 5)
}

func TestList_Get(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	list := NewListWithArray(array)
	head := list.Head().Next()
	assert.Equal(t, head.Data(), 1)
	head = head.Next()
	assert.Equal(t, head.Data(), 2)
	head = head.Next()
	assert.Equal(t, head.Data(), 3)
	head = head.Next()
	assert.Equal(t, head.Data(), 4)
	head = head.Next()
	assert.Equal(t, head.Data(), 5)
}

func TestList_Find(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	list := NewListWithArray(array)
	assert.Equal(t, list.Find(1, 5, list.Tail()).Data(), 1)
}

func TestList_InsertBefore(t *testing.T) {
	array := []int{5}
	list := NewListWithArray(array)
	list.InsertBefore(4, list.Tail())
	assert.Equal(t, list.Tail().Prev().Data(), 4)
}

func TestList_InsertAfter(t *testing.T) {
	array := []int{5}
	list := NewListWithArray(array)
	list.InsertAfter(list.Head(), 4)
	assert.Equal(t, list.Head().Next().Data(), 4)
}

func TestList_Copy(t *testing.T) {
	array := []int{5}
	list1 := NewListWithArray(array)
	array = []int{1, 2, 3, 4, 5}
	list2 := NewListWithArray(array)
	list2.Copy(list1.Head().Next(), 1)
	assert.Equal(t, list2.Tail().Prev().Data(), 5)
}

func TestList_Remove(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	list := NewListWithArray(array)
	list.Remove(list.Head().Next())
	assert.Equal(t, list.Head().Next().Data(), 2)
}

func TestList_Deduplicate(t *testing.T) {
	array := []int{1, 2, 3, 4, 4, 5}
	list := NewListWithArray(array)
	list.Deduplicate()
	assert.Equal(t, list.ArrayData(), []int{1, 2, 3, 4, 5})
}

func TestList_Uniquify(t *testing.T) {
	array := []int{1, 2, 3, 4, 4, 5}
	list := NewListWithArray(array)
	list.Uniquify()
	assert.Equal(t, list.ArrayData(), []int{1, 2, 3, 4, 5})
}

func TestList_Search(t *testing.T) {
	array := []int{1, 2, 3, 4, 5}
	list := NewListWithArray(array)
	assert.Equal(t, list.Search(1, 5, list.Tail()).Data(), 1)
}

