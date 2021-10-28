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
Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Level Order Traversal.
Memory Usage: 3.1 MB, less than 46.88% of Go online submissions for Binary Tree Level Order Traversal.
*/

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	dict := make(map[int]bool)

	var dfs func(root *TreeNode, current int)
	dfs = func(root *TreeNode, current int) {
		if root == nil {
			return
		}

		_, ok := dict[current]
		if !ok {
			ans = append(ans, []int{})
			dict[current] = true
		}
		if len(ans[current]) == 0 {
			ans[current] = []int{
				root.Val,
			}
		} else {
			ans[current] = append(ans[current], root.Val)
		}
		dfs(root.Left, current+1)
		dfs(root.Right, current+1)
	}

	dfs(root, 0)
	return ans
}

func levelOrder2(root *TreeNode) [][]int {
	var ans [][]int

	var dfs func(root *TreeNode, current int)
	dfs = func(root *TreeNode, current int) {
		if root == nil {
			return
		}

		if len(ans) <= current {
			ans = append(ans, []int{})
		}

		ans[current] = append(ans[current], root.Val)

		dfs(root.Left, current+1)
		dfs(root.Right, current+1)
	}

	dfs(root, 0)
	return ans
}
