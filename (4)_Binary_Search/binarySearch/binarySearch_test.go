package main

import "testing"

func TestTableBinarySearch(t *testing.T) {
	var tests = []struct {
		input1   []int
		input2   int
		expected int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
	}

	for i, test := range tests {
		if output := BinarySearch(test.input1, test.input2); output != test.expected {
			t.Errorf("Test %d failed for Binary Search (regular)", i)
		}
		if output := BinarySearchRecursive(test.input1, test.input2); output != test.expected {
			t.Errorf("Test %d failed for Binary Search (Recursive)", i)
		}
	}
}
