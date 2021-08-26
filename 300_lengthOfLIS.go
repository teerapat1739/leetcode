package leetcode

import "fmt"

type Sub struct {
	val   int
	index int
	lis   int
}

/*
Success
Runtime: 1280 ms, faster than 5.01% of Go online submissions for Longest Increasing Subsequence.
Memory Usage: 7.1 MB, less than 5.01% of Go online submissions for Longest Increasing Subsequence.

*/
func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	max := 1
	m := make(map[int]Sub)
	for i := 0; i < length; i++ {
		m[i] = Sub{nums[i], i, 1}
		// j := m[i].index
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				a := Max(m[i].lis, m[j].lis+1)
				m[i] = Sub{nums[i], i, a}
				max = Max(max, a)
				// break
			}
		}
	}
	fmt.Println(m)
	return max
}

/*
use: dynamic programing
Success
Runtime: 532 ms, faster than 5.01% of Go online submissions for Longest Increasing Subsequence.
Memory Usage: 6.4 MB, less than 5.01% of Go online submissions for Longest Increasing
*/

func lengthOfLIS2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	max := 1
	m := make(map[int]Sub)
	for i := 0; i < length; i++ {
		m[i] = Sub{nums[i], i, 1}
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				a := Max(m[i].lis, m[j].lis+1)
				m[i] = Sub{nums[i], i, a}
				if a > max {
					max = Max(max, a)
					break
				}
				max = Max(max, a)
			}
		}
	}
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
