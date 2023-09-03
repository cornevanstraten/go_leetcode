package main

import "testing"

func TestTableBuyAndSellStonks(t *testing.T) {
	var tests = []struct {
		input1   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{7, 2, 4, 3, 1}, 2},
		{[]int{7, 6, 4, 3, 10}, 7},
	}

	for i, test := range tests {
		if output := MaxProfit(test.input1); output != test.expected {
			t.Errorf("Test %d failed", i)
		}
	}
}
