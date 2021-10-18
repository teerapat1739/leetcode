package leetcode

import "math"

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
		currPrice := prices[i]
		if i == 0 {
			bestBuy = profit - currPrice
		} else {
			newBestBuy := profit - currPrice
			bestBuy = max(bestBuy, newBestBuy)
		}
		sell := bestBuy + currPrice - fee
		profit = max(profit, sell)
	}
	return profit
}

func maxProfitWithFreeII(prices []int, fee int) int {
	bestBuy := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		newBestBuy := profit - prices[i]
		bestBuy = min(bestBuy, newBestBuy)
		sell := prices[i] - fee - bestBuy
		if sell <= 0 {
			continue
		}
		profit += sell
		bestBuy = math.MaxInt16
	}

	return profit
}

/*
dp solution

case 1: I don't have stock on day i. Represent by dp[i][0], select max of the below:
	- case 1: I sold the stock today
			dp[i][0] = dp[i-1][i] + price[i]
	- case 2: I sold a stock at some previous day. Doing nothing today.
			dp[i][0] = dp[i-1][0]
case 2: I have stock on day i Represented by dp[i][1], select max of the below:
	- case 1: I bought it today
			dp[i][1] = dp[i-1][0] - price[i] - fee
	- case 2: I am carrying a pre-bought stock. Doing nothing today.
			dp[i][1] = dp[i-1][1]
*/

func maxProfitWithFreeIII(prices []int, fee int) int {
	lenPrices := len(prices)
	if lenPrices < 1 {
		return 0
	}

	dp := make([][]int, lenPrices)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0] - fee

	for i := 1; i < lenPrices; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}
	return dp[lenPrices-1][0]
}
