package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return helper(l1, l2, 0)

}

func helper(left *ListNode, right *ListNode, carry int) *ListNode {
	if left == nil && right == nil {
		if carry > 0 {
			return &ListNode{
				Val: carry,
			}
		}
		return nil
	}
	if left == nil {
		r := right.Val + carry
		return &ListNode{
			Val:  r % 10,
			Next: helper(nil, right.Next, r/10),
		}
	} else if right == nil {
		r := left.Val + carry
		return &ListNode{
			Val:  r % 10,
			Next: helper(left.Next, nil, r/10),
		}
	}

	r := left.Val + right.Val + carry
	return &ListNode{
		Val:  r % 10,
		Next: helper(left.Next, right.Next, r/10),
	}
}
func main() {
}
