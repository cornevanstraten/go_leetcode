package main

import "testing"

func TestTableValidAnagram(t *testing.T) {
	var tests = []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"yes", "sey", true},
		{"police", "lopice", true},
		{"whatsup", "icecream", false},
		{"goodsir", "foodsir", false},
	}

	for i, test := range tests {
		if output := IsAnagram(test.input1, test.input2); output != test.expected {
			t.Errorf("Test %d failed", i)
		}
	}
}
