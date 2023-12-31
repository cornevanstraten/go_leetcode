/*
CONTAINS DUPLICATE (217)
https://leetcode.com/problems/contains-duplicate/

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.



Example 1:

Input: nums = [1,2,3,1]
Output: true
Example 2:

Input: nums = [1,2,3,4]
Output: false
Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true


Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
*/

package main

import "fmt"

func ContainsDuplicate(nums []int) bool {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, hasNum := set[num]; hasNum {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}

func main() {

	test1 := []int{1, 2, 3, 1}
	test2 := []int{1, 2, 3, 4}
	test3 := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}

	fmt.Println(ContainsDuplicate(test1))
	fmt.Println(ContainsDuplicate(test2))
	fmt.Println(ContainsDuplicate(test3))
}

//note: an empty struct takes up no space,
//which is why it's a slightly better choice here
//than a boolean
