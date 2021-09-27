package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Get Maximum in Generated Array.
Memory Usage: 2.1 MB, less than 96.15% of Go online submissions for Get Maximum in Generated Array.
*/
func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	count := 1
	max := 1
	for (2 <= 2*count) && (2*count <= n) {
		dp[2*count] = dp[count]
		max = Max(max, dp[2*count])
		if (2 <= 2*count+1) && (2*count+1 <= n) {
			dp[2*count+1] = dp[count] + dp[count+1]
			max = Max(max, dp[2*count+1])
			count++
		} else {
			break
		}
	}

	return max
}
