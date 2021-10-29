package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var ans [][]int

	var travel func(root *TreeNode, level int)

	travel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(ans) <= level {
			ans = append(ans, []int{})
		}

		ans[level] = append(ans[level], root.Val)

		travel(root.Left, level+1)
		travel(root.Right, level+1)
	}

	travel(root, 0)

	temp := [][]int{}

	temp = append(temp, ans...)
	for i := len(temp) - 1; i >= 0; i-- {
		idx := len(ans) - 1 - i
		ans[idx] = temp[i]
	}

	return ans
}
