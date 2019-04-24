package main

type Vector struct {
	data []int
}

func NewVectorWithData(data []int) *Vector {
	return &Vector{}
}

func NewVector(length int, capacity int) *Vector {
	return &Vector{
		data: make([]int, length, capacity),
	}
}

func (vector *Vector) Data() []int {
	return vector.data
}

func (vector *Vector) Length() int {
	return len(vector.data)
}

func (vector *Vector) Capacity() int {
	return cap(vector.data)
}

func (vector *Vector) Expand() {
	if len(vector.data) < cap(vector.data) {
		return
	}
	newData := make([]int, vector.Length(), vector.Capacity()<<1)
	for index, value := range vector.data {
		newData[index] = value
	}
	vector.data = newData
}

func (vector *Vector) Shrink() {
	if vector.Length()<<1 < vector.Capacity() {
		newData := make([]int, vector.Length(), vector.Capacity()>>1)
		for index, value := range vector.data {
			newData[index] = value
		}
		vector.data = newData
	}
}

func (vector *Vector) Insert(index int, value int) {
	vector.Expand()
	length := vector.Length()
	vector.data = vector.data[0 : length+1]
	for i := length; i > index; i-- {
		vector.data[i] = vector.data[i-1]
	}
	vector.data[index] = value
}

func (vector *Vector) DeleteMany(low int, high int) {
	switch {
	case vector.Length() < high:
		panic("out of range")
	case high < vector.Length():
		for high < vector.Length() {
			vector.data[low] = vector.data[high]
			low++
			high++
		}
		vector.data = vector.data[0:low]
	default:
		vector.data = vector.data[0:low]
	}
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
	length := vector.Length()
	for index, value := range vector.data {
		if index > 0 {
			prefixIndex := vector.Find(value, 0, index-1)
			if prefixIndex != 0 {
				vector.DeleteOne(prefixIndex)
			}
		}
		if index < length-1 {
			suffixIndex := vector.Find(value, index+1, len(vector.data)-1)
			if suffixIndex != index+1 {
				vector.DeleteOne(suffixIndex)
			}
		}
	}
}

func (vector *Vector) Uniquify() {
	i, j := 0, 0
	for j < vector.Length() {
		if vector.data[j] != vector.data[i] {
			i++
			vector.data[i] = vector.data[j]
		}
		j++
	}
	vector.data = vector.data[0:i]
	vector.Shrink()
}

func (vector *Vector) BinSearch(value int, low int, high int) int {
	for low < high {
		middle := (low + high) >> 1
		if vector.data[middle] < value {
			high = middle
		} else if value < vector.data[middle] {
			low = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

func (vector *Vector) FibSearch(value int, low int, high int) int {
	fib := NewFib(high - low)
	for low < high {
		for high-low < fib.get() {
			fib.prev()
		}
		middle := low + fib.get() - 1
		if value < vector.data[middle] {
			high = middle
		} else if vector.data[middle] < value {
			low = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

func (vector *Vector) BubbleSort() {
	isSort := false
	last := vector.Length()
	for !isSort {
		isSort = true
		for i := 0; i < last-1; i++ {
			if vector.data[i] > vector.data[i+1] {
				vector.data[i], vector.data[i+1] = vector.data[i+1], vector.data[i]
				isSort = false
				last = i + 1
			}
		}
	}
}
