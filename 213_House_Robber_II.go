package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber II.
Memory Usage: 2.1 MB, less than 67.39% of Go online submissions for House Robber II.
*/
func robII(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// max(except first, expect last)
	return max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}

/*
2,3,2 | 2,3,2
2,3,

1, 2, 3, 1, | 1, 2, 3, 1,
1 2 4 3

// 2 2 2  9 1 1  1 1  9  1  7  1
// 2 2 4 11 5 12 6 13 21 14 28 22

0 2 4  9 1 1  1 1  9  1  7  1
0 2 4 11 5 12 12 13 21 14 28
*/
