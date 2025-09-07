package main

import "testing"

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 4, 2}, []int{2, 3, 4, 5, 8}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{10}, []int{10}},
		{[]int{}, []int{}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{7, 2, 5, 2, 7}, []int{2, 2, 5, 7, 7}},
		{[]int{-1, -5, 3, 0, 2}, []int{-5, -1, 0, 2, 3}},
		{[]int{100, 50, 50, 0, -100}, []int{-100, 0, 50, 50, 100}},
	}

	for _, tc := range tests {
		actual := bubbleSort(tc.input)
		if !slicesEqual(actual, tc.expected) {
			t.Errorf("bubbleSort(%v) = %v; expected %v", tc.input, actual, tc.expected)
		}
	}
}
