package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Runtime: 3117 ms, faster than 7.14% of Go online submissions for All Possible Full Binary Trees.
Memory Usage: 13.4 MB, less than 9.52% of Go online submissions for All Possible Full Binary Trees.
*/
func allPossibleFBT(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{&TreeNode{Val: 0}}
	}
	res := []*TreeNode{}
	for i := 1; i < n; i++ {
		left := allPossibleFBT(i)
		right := allPossibleFBT(n - i - 1)

		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: 0}
				root.Left = l
				root.Right = r
				res = append(res, root)
			}
		}
	}
	return res
}

/*
Runtime: 16 ms, faster than 88.10% of Go online submissions for All Possible Full Binary Trees.
Memory Usage: 9.4 MB, less than 57.14% of Go online submissions for All Possible Full Binary Trees.
*/
var memo = make([][]*TreeNode, 20+1)

func allPossibleFBTV2(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{&TreeNode{Val: 0}}
	}
	if memo[n] != nil {
		return memo[n]
	}

	res := []*TreeNode{}
	for i := 1; i < n; i++ {
		left := allPossibleFBTV2(i)
		right := allPossibleFBTV2(n - i - 1)

		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: 0}
				root.Left = l
				root.Right = r
				fmt.Println(n, i, l, r)
				res = append(res, root)
			}
		}
	}

	memo[n] = res
	return res
}
