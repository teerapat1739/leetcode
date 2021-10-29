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
Runtime: 0 ms, faster than 100.00% of Go online submissions for Minimum Distance Between BST Nodes.
Memory Usage: 2.5 MB, less than 73.33% of Go online submissions for Minimum Distance Between BST Nodes.
*/
func minDiffInBST(root *TreeNode) int {
	var arr []int

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}
		arr = append(arr, root.Val)
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)

	minDiff := math.MaxInt64

	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < minDiff {
			minDiff = arr[i] - arr[i-1]
		}
	}
	return minDiff
}
