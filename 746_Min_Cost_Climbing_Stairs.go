package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Min Cost Climbing Stairs.
Memory Usage: 3.2 MB, less than 30.57% of Go online submissions for Min Cost Climbing Stairs.
*/
func minCostClimbingStairs(cost []int) int {
	//
	dp := make([]int, len(cost)+1)
	dp[0] = cost[0]
	dp[1] = cost[1]
	if len(cost) == 2 {
		return min(dp[0], dp[1])
	}

	dp[2] = min(dp[0], dp[1]) + cost[2]

	for i := 3; i < len(dp)-1; i++ {
		dp[i] += min(dp[i-1], dp[i-2]) + cost[i]
	}

	return min(dp[len(cost)-1], dp[len(cost)-2])
}
