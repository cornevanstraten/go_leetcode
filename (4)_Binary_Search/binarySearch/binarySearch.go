/*

Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.

You must write an algorithm with O(log n) runtime complexity.



Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4
Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1


Constraints:

1 <= nums.length <= 104
-104 < nums[i], target < 104
All the integers in nums are unique.
nums is sorted in ascending order.

*/

package main

import "fmt"

func BinarySearch(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		m := (l + r) / 2
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}

func BinarySearchRecursive(nums []int, target int) int {
	var recursiveSearch func(nums []int, l int, r int) int

	recursiveSearch = func(nums []int, l int, r int) int {
		if l <= r {
			m := (l + r) / 2
			if nums[m] > target {
				return recursiveSearch(nums, l, m-1)
			} else if nums[m] < target {
				return recursiveSearch(nums, m+1, r)
			} else {
				return m
			}
		}
		return -1
	}
	return recursiveSearch(nums, 0, len(nums)-1)
}

func main() {
	fmt.Println(BinarySearch([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println("Hello, binarySearch!")
}
