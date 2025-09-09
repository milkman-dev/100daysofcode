package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7, 6},
		{[]int{2, 4, 6, 8, 10, 12, 14, 16}, 12, 5},
		{[]int{3, 6, 9, 12, 15, 18, 21}, 4, -1},
		{[]int{-50, -20, -10, -5, 0, 5, 10, 20, 50}, -10, 2},
		{[]int{10, 100, 1000, 10000, 100000}, 10000, 3},
		{[]int{1, 4, 7, 12, 18, 19, 25, 33, 47, 59, 60, 72}, 33, 7},
		{[]int{0, 1, 1, 1, 2, 3, 4, 5}, 1, 3},
		{[]int{7, 7, 7, 7, 7, 7, 7, 7}, 7, 3},
		{[]int{-1000, -500, -100, -50, -10, -5, -1, 0, 1, 5, 10, 50, 100, 500, 1000}, 500, 13},
		{[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}, 29, 9},
		{[]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}, 89, 10},
		{[]int{1, 2, 3, 4, 5}, 42, -1},
		{[]int{}, 1, -1},
	}

	for _, tc := range tests {
		actual := binarySearch(tc.input, tc.target)
		if tc.expected != actual {
			t.Errorf("Binary search (%v, %v) = expected: %v actual: %v", tc.input, tc.target, tc.expected, actual)
		}
	}

}
