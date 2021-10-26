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
Runtime: 8 ms, faster than 65.71% of Go online submissions for Find Largest Value in Each Tree Row.
Memory Usage: 6.3 MB, less than 41.43% of Go online submissions for Find Largest Value in Each Tree Row.
*/
func largestValues(root *TreeNode) []int {
	var ans []int
	dictLevel := make(map[int]int)
	level := 1
	maxLevel := 1
	maxLevel := 1
	var travel func(root *TreeNode, currentLevel int)
	travel = func(root *TreeNode, currentLevel int) {
		maxLevel = max(currentLevel, maxLevel)
		if root == nil {
			return
		}

		val, ok := dictLevel[currentLevel]
		if ok {
			dictLevel[currentLevel] = max(val, root.Val)
		} else {
			dictLevel[currentLevel] = root.Val
		}
		travel(root.Left, currentLevel+1)
		travel(root.Right, currentLevel+1)

	}
	travel(root, level)

	for i := 1; i <= maxLevel; i++ {
		if val, ok := dictLevel[i]; ok {
			ans = append(ans, val)
		}
	}

	return ans
}
