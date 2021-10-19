package leetcode

/*
dp[len][2]
case 1: We have a stock on day i, dp[i][1] max of the below:
	- I bought it today
		dp[i-2][0] - prices[i]
	- I am carry forwarding
		dp[i-1][1]

case 2: We have no stock on day i, dp[i][0] max of the below:
	- I sold it today
		dp[i-1][1] + prices[i]
	- I am carry forwarding, doing nothing
		dp[i-1][0]

*/
func maxProfitWithCoolDown(prices []int) int {
	lenPrices := len(prices)

	if lenPrices <= 1 {
		return 0
	}

	if lenPrices == 2 && prices[1] > prices[0] {
		return prices[1] - prices[0]
	} else if lenPrices == 2 && prices[0] > prices[1] {
		return 0
	}

	dp := make([][]int, lenPrices)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[1][0] = max(dp[0][0], dp[0][1]+prices[1])
	dp[1][1] = max(dp[0][1], dp[0][0]-prices[1])

	for i := 2; i < lenPrices; i++ {
		dp[i][1] = max(dp[i-2][0]-prices[i], dp[i-1][1])
		dp[i][0] = max(dp[i-1][1]+prices[i], dp[i-1][0])
	}
	return dp[lenPrices-1][0]
}
