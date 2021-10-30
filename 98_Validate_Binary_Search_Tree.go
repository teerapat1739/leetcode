package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
Runtime: 4 ms, faster than 93.94% of Go online submissions for Validate Binary Search Tree.
Memory Usage: 5.2 MB, less than 75.25% of Go online submissions for Validate Binary Search Tree.
*/
func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return validate(root.Left, min, root.Val) && validate(root.Right, root.Val, max)
}
