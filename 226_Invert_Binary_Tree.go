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
Runtime: 0 ms, faster than 100.00% of Go online submissions for Invert Binary Tree.
Memory Usage: 2.2 MB, less than 57.03% of Go online submissions for Invert Binary Tree.
*/
func invertTree(root *TreeNode) *TreeNode {
	return invertNode(root)
}

func invertNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	nextRightNode := root.Right
	nextLeftNode := root.Left
	temp := nextRightNode
	nextRightNode = nextLeftNode
	nextLeftNode = temp

	root.Left = invertNode(nextLeftNode)
	root.Right = invertNode(nextRightNode)

	return root
}
