package leetcode

func rob(nums []int) int {
	dp := make([]int, len(nums))
	result := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = nums[i]
		// find max before current position - 1
		// start pos must ge i - 2
		maxPos := i - 2
		maxVal := 0
		// limit not loop more than 5 times
		limitFindMax := 5
		for maxPos >= 0 && limitFindMax > 0 {
			maxVal = max(maxVal, dp[maxPos])
			maxPos--
			limitFindMax--
		}

		dp[i] += maxVal
		result = max(dp[i], result)
	}

	return result
}

// 9 1 9 100 1 4
// 9 1 18 109 19 113

// 2 2 2  9 1 1  1 1  9  1  7  1
// 2 2 4 11 5 12 6 13 21 14 28 22

// 1 2 3 1
// 1 2 4 3

// 2 7 9 3 1
// 2 7 11 10 12
