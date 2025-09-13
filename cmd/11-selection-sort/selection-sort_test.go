package main

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		input []int
		expected []int
	}{
		{input: []int{}, expected: []int{}},
		{input: []int{42}, expected: []int{42}},
		{input: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{input: []int{3, 1, 4, 1, 5, 9, 2}, expected: []int{1, 1, 2, 3, 4, 5, 9}},
		{input: []int{-3, -1, -7, 4, 2}, expected: []int{-7, -3, -1, 2, 4}},
		{input: []int{8, 8, 8, 8}, expected: []int{8, 8, 8, 8}},
		{input: []int{0, -1, 5, -10, 3}, expected: []int{-10, -1, 0, 3, 5}},
		{input: []int{1000, 50, 700, -300, 25}, expected: []int{-300, 25, 50, 700, 1000}},
		{input: []int{4, 2, 2, 8, 5, 6, 2, 8}, expected: []int{2, 2, 2, 4, 5, 6, 8, 8}},
	}

	for _, tc := range tests {
		actual := selectionSort(tc.input)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("Failed:\n expected: %v\n actual: %v", tc.expected, actual)
		}
	}

}
