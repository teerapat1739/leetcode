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
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Postorder Traversal.
Memory Usage: 2.1 MB, less than 99.74% of Go online submissions for Binary Tree Postorder Traversal.
*/
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ans []int
	// var stack []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}

		if root.Right != nil {
			dfs(root.Right)
		}
		ans = append(ans, root.Val)
	}
	dfs(root)
	return ans
}
