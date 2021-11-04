package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Good Pairs.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Number of Good Pairs.
*/

func numIdenticalPairs(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				count++
			}
		}
	}
	return count
}
