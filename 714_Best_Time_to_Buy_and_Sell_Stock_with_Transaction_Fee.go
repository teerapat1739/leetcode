package leetcode

/*
Success
Details
Runtime: 88 ms, faster than 67.50% of Go online submissions for Best Time to Buy and Sell Stock with Transaction Fee.
Memory Usage: 8.3 MB, less than 40.00% of Go online submissions for Best Time to Buy and Sell Stock with Transaction Fee.
*/
func maxProfitWithFree(prices []int, fee int) int {
	bestBuy := 0
	profit := 0
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			bestBuy = -prices[i]
		} else {
			newBestBuy := -prices[i] + profit
			bestBuy = max(bestBuy, newBestBuy)
		}
		profit = max(profit, bestBuy+prices[i]-fee)
	}
	return profit
}
