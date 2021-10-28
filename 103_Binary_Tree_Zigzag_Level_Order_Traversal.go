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
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
Memory Usage: 2.9 MB, less than 17.71% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
*/

func zigzagLevelOrder(root *TreeNode) [][]int {
	var ans [][]int
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(ans) <= level {
			ans = append(ans, []int{})
		}

		if level%2 == 0 {
			ans[level] = append(ans[level], root.Val)
		} else {
			ans[level] = append([]int{root.Val}, ans[level]...)
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return ans
}
