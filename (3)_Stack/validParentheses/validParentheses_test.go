package main

import "testing"

func TestTableValidParentheses(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(])", false},
		{"{([[{{[{}]}}]])}", true},
	}

	for i, test := range tests {
		if output := ValidParentheses(test.input); output != test.expected {
			t.Errorf("Test %d failed", i)
		}
	}
}
