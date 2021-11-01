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
Runtime: 8 ms, faster than 92.52% of Go online submissions for Kth Smallest Element in a BST.
Memory Usage: 6.4 MB, less than 85.33% of Go online submissions for Kth Smallest Element in a BST.
*/
func kthSmallest(root *TreeNode, k int) int {
	var ans []int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if k == len(ans) {
			return
		}
		if root == nil {
			return
		}

		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}
	dfs(root)

	return ans[k-1]

}
