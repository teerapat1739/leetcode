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
 Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Inorder Traversal.
Memory Usage: 2.1 MB, less than 99.89% of Go online submissions for Binary Tree Inorder Traversal.
*/
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var ans []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			dfs(root.Left)
		}
		ans = append(ans, root.Val)

		if root.Right != nil {
			dfs(root.Right)
		}
	}
	dfs(root)
	return ans
}
