package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths II.
Memory Usage: 2.6 MB, less than 76.44% of Go online submissions for Unique Paths II.
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[i]))
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if obstacleGrid[i][j] == 1 {
				continue
			}
			if i == 0 && j == 0 {
				dp[i][j] = 1
			}
			if i != 0 && j != 0 {
				dp[i][j] += dp[i][j-1] + dp[i-1][j]
				continue
			}
			if i == 0 && j != 0 {
				dp[i][j] += dp[i][j-1]
				continue
			}
			if i != 0 && j == 0 {
				dp[i][j] += dp[i-1][j]
				continue
			}
		}
	}

	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
