package main

import "math"

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	dp := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		dp[i] = math.MaxInt32
		for _, denom := range {
			if i >= denom {
				dp[i] = min(dp[i], dp[i-denom])
			}                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              
		}
	}
	if dp[n] == math.MaxInt32 {
			return -1
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
