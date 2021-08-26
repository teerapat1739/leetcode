package leetcode

import "math"

/*
Success
Details:
Runtime: 120 ms, faster than 75.22% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 8.8 MB, less than 50.05% of Go online submissions for Best Time to Buy and Sell Stock.
*/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	bestBuy := math.MaxInt64
	bestValue := 0
	for i := 0; i < len(prices); i++ {
		if bestBuy > prices[i] {
			bestBuy = prices[i]
		}
		if prices[i]-bestBuy > bestValue {
			bestValue = prices[i] - bestBuy
		}
	}
	return bestValue
}
