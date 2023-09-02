package main

import "testing"

func TestTableTwoSum(t *testing.T) {
	var tests = []struct {
		input1   []int
		input2   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for i, test := range tests {
		output := TwoSum(test.input1, test.input2)

		if !EqualSlices(output, test.expected) {
			t.Errorf("Test %d failed", i)
		}
	}
}
