package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.
Memory Usage: 2.2 MB, less than 24.67% of Go online submissions for Unique Paths.
*/
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// assign 1 in row m = 0 and n = 0
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] += dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
