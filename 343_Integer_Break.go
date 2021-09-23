package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Integer Break.
Memory Usage: 2 MB, less than 84.38% of Go online submissions for Integer Break.
*/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	//dp := make(map[int][]int, n)

	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = Max(dp[i], Max(j, dp[j])*Max(i-j, dp[i-j]))
		}
	}

	return dp[n]
}
