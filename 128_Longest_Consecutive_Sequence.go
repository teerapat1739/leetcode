package leetcode

import "sort"

/*
Success
Details
Runtime: 40 ms, faster than 78.89% of Go online submissions for Longest Consecutive Sequence.
Memory Usage: 7.3 MB, less than 89.27% of Go online submissions for Longest Consecutive Sequence.

*/
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	max := 0
	dp := make([]int, len(nums)+1)
	for i := 1; i < len(nums); i++ {
		prev := nums[i-1]
		curr := nums[i]
		if curr-prev == 1 {
			dp[i] = dp[i-1] + 1
		} else if curr-prev == 0 {
			dp[i] = dp[i-1]

		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max + 1
}
