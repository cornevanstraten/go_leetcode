package main

import "testing"

func TestTableContainsDuplicate(t *testing.T) {
	var tests = []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for i, test := range tests {
		if output := ContainsDuplicate(test.input); output != test.expected {
			t.Errorf("Test %d failed", i)
		}
	}
}
