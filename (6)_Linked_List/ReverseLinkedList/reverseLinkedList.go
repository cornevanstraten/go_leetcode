/*Given the head of a singly linked list, reverse the list, and return the reversed list.



Example 1:


Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
Example 2:


Input: head = [1,2]
Output: [2,1]
Example 3:

Input: head = []
Output: []


Constraints:

The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000


Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	//traverse to the end of the list
	for curr != nil {
		next := curr.Next //tmp store next
		curr.Next = prev  //reverse the link
		prev = curr       //move up prev
		curr = next       //move up curr
	}

	return prev //cur == nil; so prev = head

}

//RECURSIVE OPTION (slower, and slightly less memory efficient)

func ReverseListRecursive(head *ListNode) *ListNode {
	curr := head
	if curr == nil || curr.Next == nil { //base case: tail found
		return curr
	}

	newHead := ReverseListRecursive(curr.Next) //find tail, turn into new head
	curr.Next.Next = curr                      //reverse the link
	curr.Next = nil                            //create new tail
	return newHead
}

func main() {
	fmt.Println("Hello, reverse linked list")
}
