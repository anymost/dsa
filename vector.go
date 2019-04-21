package main

import "fmt"

type Vector struct {
	data     []int
	length   int
	capacity int
}

func NewVector(length int, capacity int) *Vector {
	return &Vector{
		data:     make([]int, length, capacity),
		length:   length,
		capacity: capacity,
	}
}

func (vector *Vector) Expand() {
	if vector.length < vector.capacity {
		return
	}
	vector.capacity = vector.capacity << 1
	newData := make([]int, vector.length, vector.capacity)
	for index, value := range vector.data {
		newData[index] = value
	}
	vector.data = newData
}

func (vector *Vector) Shrink() {
	if vector.length<<1 < vector.capacity {
		vector.capacity = vector.capacity >> 1
		newData := make([]int, vector.length, vector.capacity)
		for index, value := range vector.data {
			newData[index] = value
		}
		vector.data = newData
	}
}

func (vector *Vector) Insert(index int, value int) {
	vector.Expand()
	vector.data = vector.data[0 : vector.length+1]
	for i := vector.length; vector.length > index; i-- {
		vector.data[i] = vector.data[i-1]
	}
	vector.data[index] = value
	vector.length = vector.length + 1
}

func (vector *Vector) DeleteMany(low int, high int) {
	for high < vector.length {
		vector.data[low] = vector.data[high]
		low++
		high++
	}
	vector.length = low
	vector.data = vector.data[0:vector.length]
	vector.Shrink()
}

func (vector *Vector) DeleteOne(index int) {
	vector.DeleteMany(index, index+1)
}

func (vector *Vector) Find(value int, low int, high int) int {
	for ; high > low && vector.data[high] != value; high-- {
	}
	return high
}

func (vector *Vector) Deduplicate() {
	for index, value := range vector.data {
		prefixIndex := vector.Find(value, 0, index-1)
		if prefixIndex != 0 {
			vector.DeleteOne(prefixIndex)
		}
		suffixIndex := vector.Find(value, index+1, vector.length-1)
		if suffixIndex != index+1 {
			vector.DeleteOne(suffixIndex)
		}
	}
}

func (vector *Vector) Uniquify() {
	i, j := 0, 0
	for j < vector.length {
		if vector.data[j] != vector.data[i] {
			i += 1
			vector.data[i] = vector.data[j]
			j++
		}
	}
	vector.length = i
	vector.Shrink()
}

func (vector *Vector) BinSearch(value int, low int, high int) int {
	for low < high {
		middle := (low + high) >> 1
		if vector.data[middle] < value {
			low = middle
		} else if value < vector.data[middle] {
			high = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

func main() {
	vec := NewVector(0, 4)
	vec.Insert(0, 1)
	vec.Insert(1, 2)
	vec.Insert(2, 3)
	vec.Insert(3, 4)
	vec.Insert(4, 4)
	vec.Insert(5, 5)
	vec.Insert(6, 6)
	index := vec.BinSearch(10, 0, 7)
	fmt.Println(index)

}
