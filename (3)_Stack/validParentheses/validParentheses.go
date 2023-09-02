/*
Valid Parentheses
Leetcode question 20
https://leetcode.com/problems/valid-parentheses/

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.

*/

package main

import "fmt"

func ValidParentheses(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	var stack []rune

	m := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}

	for _, val := range s {
		if _, ok := m[val]; ok {
			stack = append(stack, val)
		} else {
			idx := len(stack) - 1
			if idx < 0 || m[stack[idx]] != val {
				return false
			}
			stack = stack[:idx]
		}
	}

	length := len(stack)
	return length == 0
}

func main() {
	fmt.Println(ValidParentheses("(({{}}))"))
	fmt.Println("Hello, validParentheses")
}
