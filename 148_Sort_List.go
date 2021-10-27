package leetcode

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
Memory Usage: 2.1 MB, less than 44.64% of Go online submissions for Sort Colors.
*/
func sortList(head *ListNode) *ListNode {
	arr := []int{}
	dict := make(map[int]*ListNode)

	curr := head

	for curr != nil {
		val := curr.Val
		arr = append(arr, val)
		next := curr.Next
		dict[val] = next
		curr = next

	}

	sort.Ints(arr)

	ans := &ListNode{}
	for i := len(arr) - 1; i >= 0; i-- {
		ans.Val = arr[i]
		temp := ans
		ans = &ListNode{
			Next: temp,
		}
	}

	return ans.Next
}
