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
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
Memory Usage: 2.1 MB, less than 99.63% of Go online submissions for Binary Tree Preorder Traversal.

*/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ans []int

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		ans = append(ans, root.Val)
		if root.Left != nil {
			dfs(root.Left)
		}
		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return ans
}
