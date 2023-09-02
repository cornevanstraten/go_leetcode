package main

import "testing"

func TestTableIsPalindrome(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for i, test := range tests {
		if output := IsPalindrome(test.input); output != test.expected {
			t.Errorf("Test %d failed", i)
		}
	}
}
