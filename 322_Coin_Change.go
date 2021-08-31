package leetcode

import "math"

// *** fail case 7
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	memo := make(map[int]int)
	memo[0] = 0
	ans := recursive(coins, amount, math.MaxInt32, memo)
	if ans == 0 {
		return -1
	}
	return ans
}

func recursive(coins []int, amount, m int, memo map[int]int) int {
	if amount == 0 {
		return 0
	}

	v, ok := memo[amount]
	if ok {
		return v
	}
	for _, coin := range coins {
		if amount-coin >= 0 {
			m = min(m, recursive(coins, amount-coin, m, memo)+1)
			memo[amount] = m
		}
	}
	if m == math.MaxInt32 {
		return -1
	}

	return memo[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
Success
Details
Runtime: 4 ms, faster than 98.91% of Go online submissions for Coin Change.
Memory Usage: 6.5 MB, less than 97.59% of Go online submissions for Coin Change.
Next challenges:
*/

func coinChange2(coins []int, amount int) int {
	memo := make([]int, amount+1)
	memo[0] = 0
	for i := 1; i <= amount; i++ {
		memo[i] = math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 {
				memo[i] = min(memo[i], memo[i-coin]+1)
			}
		}
	}
	if memo[amount] == math.MaxInt32 {
		return -1
	}
	return memo[amount]
}
