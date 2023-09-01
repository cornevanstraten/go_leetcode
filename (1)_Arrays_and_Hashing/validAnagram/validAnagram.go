/*
Valid Anagram
Leetcode question 242
https://leetcode.com/problems/valid-anagram/description/

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false


Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.


Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?

*/

package main

import "fmt"

func IsAnagram(s string, t string) bool {
	//check for the same number of characters
	if len(s) != len(t) {
		return false
	}

	m := make(map[rune]int)

	//fill a map with characters and respective frequencies for the first string
	for _, c := range s {
		m[c]++
	}
	//decrement the for the characters of the second string
	for _, c := range t {
		m[c]--
		//if a character is ever decremented more than it was initially incremented, return false
		if m[c] < 0 {
			return false
		}

	}
	return true
}

func main() {
	fmt.Println("Hello, anagram!")
}
