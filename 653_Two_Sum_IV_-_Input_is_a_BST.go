package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
Runtime: 28 ms, faster than 46.15% of Go online submissions for Two Sum IV - Input is a BST.z
Memory Usage: 7.5 MB, less than 83.92% of Go online submissions for Two Sum IV - Input is a BST.
*/
func findTarget(root *TreeNode, k int) bool {
	var dfs func(root *TreeNode)
	ans := []int{}
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		ans = append(ans, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	fmt.Println(ans)

	start := 0
	end := len(ans) - 1

	for start < end {
		cur := ans[start] + ans[end]
		if cur == k {
			return true
		} else if cur < k {
			start++
		} else {
			end--
		}
	}
	return false
}
