package algoexpert

import "sort"

func NonConstructibleChange(coins []int) int {
	// Write your code here.
	sort.Ints(coins)
	currChange := 0
	for _, coin := range coins {
		if coin > currChange+1 {
			return currChange + 1
		}
		currChange += coin
	}

	return currChange + 1
}
