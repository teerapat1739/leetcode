package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Two Sorted Lists.
Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Merge Two Sorted Lists.
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans *ListNode

	currentNode := ans
	for l1 != nil || l2 != nil {
		if l1 == nil {
			currentNode.Next = l2
			break
		}

		if l2 == nil {
			currentNode.Next = l1
			break
		}

		if l1.Val < l2.Val {
			currentNode.Next = l1.Next
			l1 = l1.Next
		} else {
			currentNode.Next = l2.Next
			l2 = l2.Next
		}

		currentNode = currentNode.Next
	}
	return ans
}
