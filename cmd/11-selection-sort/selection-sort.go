package main

func selectionSort(arr []int) []int {
	for i := range len(arr) {
		smallestIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[smallestIdx] {
				smallestIdx = j
			}
		}
		arr[i], arr[smallestIdx] = arr[smallestIdx], arr[i]
	}
	return arr
}
