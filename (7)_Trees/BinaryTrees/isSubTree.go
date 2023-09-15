/* Leetcode 572: Subtree of another tree

Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values of subRoot and false otherwise.

A subtree of a binary tree tree is a tree that consists of a node in tree and all of this node's descendants. The tree tree could also be considered as a subtree of itself.



Example 1:


Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
Example 2:


Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false


Constraints:

The number of nodes in the root tree is in the range [1, 2000].
The number of nodes in the subRoot tree is in the range [1, 1000].
-104 <= root.val <= 104
-104 <= subRoot.val <= 104
*/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	//same as sameTree with a twist
	//compare roots
	if isSameTree(root, subRoot) {
		return true
	}
	//if not true, compare left and right until you hit the base case:
	if root == nil {
		return false
	}
	left := isSubtree(root.Left, subRoot)
	right := isSubtree(root.Right, subRoot)

	if left || right {
		return true
	}

	return false
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	//base case 0: both nil
	if p == nil && q == nil {
		return true
	}

	//base case 1: unequality
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	//recurse
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
