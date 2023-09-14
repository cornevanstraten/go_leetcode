/*Diameter of Binary Tree
Leetcode 543
https://leetcode.com/problems/diameter-of-binary-tree/

Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.



Example 1:


Input: root = [1,2,3,4,5]
Output: 3
Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].
Example 2:

Input: root = [1,2]
Output: 1


Constraints:

The number of nodes in the tree is in the range [1, 104].
-100 <= Node.val <= 100
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	dfs(root, &maxDiameter)
	return maxDiameter
}

func dfs(node *TreeNode, maxDiameter *int) int {
	if node == nil {
		return 0
	}

	left := dfs(node.Left, maxDiameter)
	right := dfs(node.Right, maxDiameter)
	if left+right > *maxDiameter {
		*maxDiameter = left + right
	}

	if left > right {
		return left + 1
	}
	return right + 1
}

//this recursive function calculates the diameter at every node and returns it
//it drills down with depth first search (DFS) till it reaches nil
//in which case it returns 0
//it returns 1 for every node up
//every iteration, we check if the local diameter (left+right) beats the current
//maxDiameter; if so it gets updated.
