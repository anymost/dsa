package main

func bubbleSort(array []int) []int {
	sort := false
	for !sort {
		sort = true
		for i := 0; i < len(array) - 1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				sort = false
			}
		}
	}
	return array
}
