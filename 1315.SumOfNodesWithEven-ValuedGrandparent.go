package leetcode

/*
Runtime: 28 ms, faster than 84.38% of Go online submissions for Sum of Nodes with Even-Valued Grandparent.
Memory Usage: 8 MB, less than 93.75% of Go online submissions for Sum of Nodes with Even-Valued Grandparent.
*/
func sumEvenGrandparent(root *TreeNode) int {
	return dfs(root, nil, nil)
}

func dfs(current, parent, grandParent *TreeNode) int {
	if current == nil {
		return 0
	}
	sum := 0
	if grandParent != nil && grandParent.Val%2 == 0 {
		sum += current.Val
	}
	sum += dfs(current.Left, current, parent)
	sum += dfs(current.Right, current, parent)
	return sum
}
