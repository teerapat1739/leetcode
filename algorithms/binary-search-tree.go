package algorithms

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
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

func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}
	if n.Value == k {
		return true
	} else if n.Value < k {
		return n.Right.Search(k)
	} else {
		return n.Left.Search(k)
	}
}

func DemoBinarySerchTree() {
	tree := &Node{Value: 100}
	fmt.Println(tree)
	//tree.Insert(200)
	//tree.Insert(300)
	//fmt.Println(tree)

	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(76))
}
