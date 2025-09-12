package main

import "fmt"

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			j--
		}

	}

	return arr
}

func main() {
	test := []int{5, 2, 7, 3, 4, 1}

	fmt.Println(test)
	fmt.Println(insertionSort(test))
}
