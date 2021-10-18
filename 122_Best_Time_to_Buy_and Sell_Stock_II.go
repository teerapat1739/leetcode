package leetcode

/*
Success
Details
Runtime: 9 ms, faster than 8.29% of Go online submissions for Best Time to Buy and Sell Stock II.
Memory Usage: 3.3 MB, less than 9.01% of Go online submissions for Best Time to Buy and Sell Stock II.

*/
func maxProfitII(prices []int) int {
	bestValue := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			bestValue += prices[i] - prices[i-1]
		}
	}
	return bestValue
}

/*
[7,1,5,3,6,4]


*/
