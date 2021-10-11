package leetcode

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	a, b := len(grid), len(grid[0])
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			// case i == 0 j == 0, only plus current position
			if i == 0 && j == 0 {
				dp[i][j] += grid[i][j]
				continue
			}
			//case i == 0, plus only left side current position
			if i == 0 && j != 0 {
				dp[i][j] += dp[i][j-1] + grid[i][j]
				continue
			}
			// case i != 0 && j == 0, current position plus only top side

			if i != 0 && j == 0 {
				dp[i][j] += dp[i-1][j] + grid[i][j]
				continue
			}

			if i != 0 && j != 0 {
				dp[i][j] += grid[i][j] + min(dp[i][j-1], dp[i-1][j])
			}

		}
	}
	return dp[a-1][b-1]
}
