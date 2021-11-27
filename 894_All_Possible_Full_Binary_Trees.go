package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 3117 ms, faster than 7.14% of Go online submissions for All Possible Full Binary Trees.
Memory Usage: 13.4 MB, less than 9.52% of Go online submissions for All Possible Full Binary Trees.
*/
func allPossibleFBT(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{&TreeNode{Val: 0}}
	}
	res := []*TreeNode{}
	for i := 1; i < n; i++ {
		left := allPossibleFBT(i)
		right := allPossibleFBT(n - i - 1)

		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: 0}
				root.Left = l
				root.Right = r
				res = append(res, root)
			}
		}
	}
	return res
}
