package main

func bubbleSort(array []int) []int {
	isSort := false
	for !isSort {
		isSort = true
		for i := 0; i < len(array) - 1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isSort = false
			}
		}
	}
	return array
}

