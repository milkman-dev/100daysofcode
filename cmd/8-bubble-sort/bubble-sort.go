package main

func bubbleSort(arr []int) []int {
	end := len(arr) - 1
	swapping := true

	for swapping {
		swapping = false
		for i := 0; i < end; i++ {
			if arr[i] > arr[end] {
				arr[i], arr[end] = arr[end], arr[i]
			} else {
				swapping = true
			}
		}
		end--
	}

	return arr
}

func main() {
}
