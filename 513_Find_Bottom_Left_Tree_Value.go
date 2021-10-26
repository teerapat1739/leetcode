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
Runtime: 4 ms, faster than 92.16% of Go online submissions for Find Bottom Left Tree Value.
Memory Usage: 5.5 MB, less than 60.78% of Go online submissions for Find Bottom Left Tree Value.
*/

func findBottomLeftValue(root *TreeNode) int {
	var max, res int
	var dfs func(root *TreeNode, current int)
	dfs = func(root *TreeNode, current int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			if current > max {
				max = current
				res = root.Val
			}
		} else {
			dfs(root.Left, current+1)
			dfs(root.Right, current+1)
		}
	}

	dfs(root, 1)
	return res
}
