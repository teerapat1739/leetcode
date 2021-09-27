package leetcode

type NodeX struct {
	Value int
	Index int
	Left  *NodeX
	Right *NodeX
}

func (n *NodeX) InsertX(val, idx int) {
	if n.Value < val {
		if n.Right == nil {
			n.Right = &NodeX{Value: val, Index: idx}
		} else {
			n.Right.InsertX(val, idx)
		}
	} else if n.Value > val {
		if n.Left == nil {
			n.Left = &NodeX{Value: val, Index: idx}
		} else {
			n.Left.InsertX(val, idx)
		}
	}
}

func (n *NodeX) SearchIdx(val int) int {
	if n == nil {
		return -1
	}
	if n.Value == val {
		return n.Index
	} else if n.Value < val {
		return n.Right.SearchIdx(val)
	} else {
		return n.Left.SearchIdx(val)
	}
}

func search(nums []int, target int) int {

	tree := &NodeX{Value: nums[0], Index: 0}
	for i := 1; i < len(nums); i++ {
		tree.InsertX(nums[i], i)
	}

	return tree.SearchIdx(target)
}
