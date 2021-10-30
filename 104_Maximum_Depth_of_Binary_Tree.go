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
Runtime: 4 ms, faster than 87.43% of Go online submissions for Maximum Depth of Binary Tree.
Memory Usage: 4.4 MB, less than 62.62% of Go online submissions for Maximum Depth of Binary Tree.
*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dfs func(root *TreeNode, level int) int
	dfs = func(root *TreeNode, level int) int {
		if root == nil {
			return level
		}
		return max(dfs(root.Left, level+1), dfs(root.Right, level+1))
	}
	return dfs(root, 0)
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
