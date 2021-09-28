package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for N-th Tribonacci Number.
Memory Usage: 2 MB, less than 52.41% of Go online submissions for N-th Tribonacci Number.
*/
func tribonacci(n int) int {
	if n == 0 {
		return n
	}

	if n < 3 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = dp[i-3] + dp[i-2] + dp[i-1]
	}
	return dp[n]
}
