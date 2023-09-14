/*
Leetcode question 110: Balanced Binary Tree
https://leetcode.com/problems/balanced-binary-tree/

Given a binary tree, determine if it is
height-balanced
.



Example 1:


Input: root = [3,9,20,null,null,15,7]
Output: true
Example 2:


Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
Example 3:

Input: root = []
Output: true


Constraints:

The number of nodes in the tree is in the range [0, 5000].
-104 <= Node.val <= 104

*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	balanced := true
	maxDepth(root, &balanced)
	return balanced
}

func maxDepth(root *TreeNode, balanced *bool) int {
	if root == nil || *balanced == false {
		return 0
	}

	//recurse and add 1 for every layer
	maxDepthLeft := maxDepth(root.Left, balanced) + 1
	maxDepthRight := maxDepth(root.Right, balanced) + 1
	diff := maxDepthLeft - maxDepthRight

	if diff > 1 || diff < -1 {
		*balanced = false
	}

	if maxDepthLeft > maxDepthRight {
		return maxDepthLeft
	}
	return maxDepthRight
}
