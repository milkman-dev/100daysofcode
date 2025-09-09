package main

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for high >= low {
		median := low + (high-low)/2
		if arr[median] == target {
			return median
		}
		if arr[median] > target {
			high = median - 1
		} else {
			low = median + 1
		}
	}
	return -1
}

func main() {
}
