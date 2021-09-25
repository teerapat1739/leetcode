package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Fibonacci Number.
Memory Usage: 2 MB, less than 62.36% of Go online submissions for Fibonacci Number.

*/
func fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
