package leetcode

import "sort"

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Array by Increasing Frequency.
Memory Usage: 3.1 MB, less than 76.19% of Go online submissions for Sort Array by Increasing Frequency.
*/
func frequencySort1636(nums []int) []int {
	dictNum := make(map[int]int)
	for _, num := range nums {
		dictNum[num] += 1
	}

	sort.Slice(nums, func(i, j int) bool {
		if dictNum[nums[i]] == dictNum[nums[j]] {
			return nums[j] < nums[i]
		}

		return dictNum[nums[i]] < dictNum[nums[j]]
	})

	return nums
}
