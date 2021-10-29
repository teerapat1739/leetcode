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
Runtime: 0 ms, faster than 100.00% of Go online submissions for Flatten Binary Tree to Linked List.
Memory Usage: 3 MB, less than 31.63% of Go online submissions for Flatten Binary Tree to Linked List.
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// tempRoot := root
	var ans []int

	var travel func(root *TreeNode)

	travel = func(root *TreeNode) {

		ans = append(ans, root.Val)
		if root.Left != nil {
			travel(root.Left)
		}
		if root.Right != nil {
			travel(root.Right)
		}
	}

	travel(root)

	root.Left = nil
	root.Right = nil
	node := root
	for i := 1; i < len(ans); i++ {
		node.Right = &TreeNode{
			Right: nil,
			Val:   ans[i],
		}
		node = node.Right
	}
}
