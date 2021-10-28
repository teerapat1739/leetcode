package leetcode

func sortedSquares(nums []int) []int {

	start, end := 0, len(nums)-1
	dp := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		powStart := nums[start] * nums[start]
		powEnd := nums[end] * nums[end]
		if powStart > powEnd {
			dp[i] = powStart
			start++
		} else {
			dp[i] = powEnd
			end--
		}
	}

	return dp
}
