package leetcode

import (
	"math"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Convert Binary Number in a Linked List to Integer.
Memory Usage: 2.1 MB, less than 30.77% of Go online submissions for Convert Binary Number in a Linked List to Integer.
*/
func getDecimalValue(head *ListNode) int {
	total := 0
	array := []int{}
	for head != nil {
		array = append(array, head.Val)
		head = head.Next
	}

	lenArr := len(array)
	idx := 0
	for i := lenArr - 1; i >= 0; i-- {
		total += array[idx] * int((math.Pow(2, float64(i))))
		idx++
	}
	return total
}
