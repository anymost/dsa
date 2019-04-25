package main

const (
	INIT_VECTOR_CAPACITY = 4
)

type Vector struct {
	data []int
}

func NewVectorWithData(data []int) *Vector {
	return &Vector{data}
}

func NewVector(length, capacity int) *Vector {
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
	var newData []int
	if vector.Capacity() != 0 {
		newData = make([]int, vector.Length(), vector.Capacity()<<1)
	} else {
		newData = make([]int, vector.Length(), INIT_VECTOR_CAPACITY)
	}
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

func (vector *Vector) Insert(index, value int) {
	vector.Expand()
	length := vector.Length()
	vector.data = vector.data[0 : length+1]
	for i := length; i > index; i-- {
		vector.data[i] = vector.data[i-1]
	}
	vector.data[index] = value
}

func (vector *Vector) DeleteMany(low, high int) {
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

func (vector *Vector) Find(value, low, high int) int {
	for high--; vector.Data()[high] != value && high > low; high-- {
	}
	return high
}

func (vector *Vector) Duplicate() {
	for index := vector.Length() - 1; index > -1; index-- {
		value := vector.Data()[index]
		prefixIndex := vector.Find(value, 0, index-1)
		if prefixIndex != 0 {
			vector.DeleteOne(prefixIndex)
		}

		suffixIndex := vector.Find(value, index+1, vector.Length()-1)
		if suffixIndex != index+1 {
			vector.DeleteOne(suffixIndex)
		}
	}
}

func (vector *Vector) Unique() {
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

func (vector *Vector) BinSearch(value, low, high int) int {
	for low < high {
		middle := (low + high) >> 1
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

func (vector *Vector) FibSearch(value, low, high int) int {
	fib := NewFib(high - low)
	for low <= high {
		for high-low < fib.Get() {
			fib.Prev()
		}
		middle := low + fib.Get() - 1
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
	isSort, index := false, vector.Length()-1
	for !isSort {
		isSort = true
		last := index
		for i := 0; i < last; i++ {
			if vector.data[i] > vector.data[i+1] {
				vector.data[i], vector.data[i+1] = vector.data[i+1], vector.data[i]
				isSort = false
				index = i
			}
		}
	}
}

func (vector *Vector) Merge(lo, mi, hi int) {
	length := hi - lo
	newData := make([]int, length, length)
	for i, j, k := 0, lo, mi; i < length; i++ {
		if j < mi && (k >= hi || vector.Data()[j] <= vector.Data()[k]) {
			newData[i], j = vector.Data()[j], j+1
		}
		if k < hi && (j >= mi || vector.Data()[k] < vector.Data()[j]) {
			newData[i], k = vector.Data()[k], k+1
		}
	}
	for index, value := range newData {
		vector.data[lo+index] = value
	}
}

func (vector *Vector) MergeSort(lo, hi int) {
	if hi-lo < 2 {
		return
	}
	mi := (lo + hi) >> 1
	vector.MergeSort(lo, mi)
	vector.MergeSort(mi, hi)
	vector.Merge(lo, mi, hi)
}
