package leetcode

/*
Runtime: 140 ms, faster than 83.60% of Go online submissions for Daily Temperatures.
Memory Usage: 9.8 MB, less than 59.79% of Go online submissions for Daily Temperatures.
*/
func dailyTemperatures(temperatures []int) []int {
	// use stack
	stack := []int{
		0,
	} // stack index of temperatures

	dp := make([]int, len(temperatures))
	for i := 1; i < len(temperatures); i++ {
		curVal := temperatures[i]
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < curVal {
			popVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			dp[popVal] = i - popVal
		}
		stack = append(stack, i)
	}
	return dp
}
