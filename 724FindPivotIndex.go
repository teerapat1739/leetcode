package leetcode

/*
Runtime: 12 ms, faster than 99.09% of Go online submissions for Find Pivot Index.
Memory Usage: 6.3 MB, less than 97.27% of Go online submissions for Find Pivot Index.
*/
func pivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	curr := 0
	for i := 0; i < len(nums); i++ {
		if curr == sum-curr-nums[i] {
			return i
		}
		curr += nums[i]
	}

	return -1
}
