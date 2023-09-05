/*
Merge two sorted lists
https://leetcode.com/problems/merge-two-sorted-lists/description/

You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.



Example 1:


Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.

*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	start := new(ListNode) //start with empty node
	tail := start          //set tail to empty to start growing the new list

	for list1 != nil && list2 != nil { //while there are (still) two lists to deal with
		if list1.Val < list2.Val { //point the tail to the lowest of the values
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 == nil { //if one of the lists is nil, point to the other one
		tail.Next = list2
	} else {
		tail.Next = list1
	}

	return start.Next //head is where we started
}

func mergeTwoListOriginal(list1 *ListNode, list2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode

	//if list1 or list 2 is null, return the other one
	if list1 == nil || list2 == nil {
		if list1 == nil {
			return list2
		}
		return list1
	}

	//if list1.val < list2.val; set head to list1 and ++ list1; otherwise for list2
	if list1.Val < list2.Val {
		head = list1
		list1 = list1.Next
	} else {
		head = list2
		list2 = list2.Next
	}

	//set tail to head and start growing the tail
	tail = head

	//while there is still elements in each list
	for list1 != nil && list2 != nil {
		//evaluate the elements, set the smallest to next, and update tail
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	//finally set the tail to the node from the list that is not empty
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return head

}

func main() {
	fmt.Println("Hello, Merge2Lists")
}
