package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
Runtime: 17 ms, faster than 49.34% of Go online submissions for Count Complete Tree Nodes.
Memory Usage: 7.5 MB, less than 38.82% of Go online submissions for Count Complete Tree Nodes.
*/
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := countNodes(root.Left)
	right := countNodes(root.Right)
	return left + right + 1
}
