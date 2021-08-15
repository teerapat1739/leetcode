package leetcode

type Node struct {
	Value uint8
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k uint8) {
	if n.Value < k {
		if n.Right == nil {
			n.Right = &Node{Value: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Value > k {
		if n.Left == nil {
			n.Left = &Node{Value: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) SearchMin() uint8 {

	if n.Left == nil {
		return n.Value
	}
	return n.Left.SearchMin()
}

func nextGreatestLetter(letters []byte, target byte) byte {
	tree := &Node{Value: target}
	for _, v := range letters {
		tree.Insert(v)
	}

	if tree.Right != nil {
		return tree.Right.SearchMin()
	} else {
		return tree.Left.SearchMin()
	}
}
