package leetcode

import "math/rand"

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Colors.
Memory Usage: 2.1 MB, less than 44.64% of Go online submissions for Sort Colors.
*/

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	left, right := 0, len(nums)-1

	pivot := rand.Int() % len(nums)

	nums[pivot], nums[right] = nums[right], nums[pivot]

	for i, _ := range nums {
		if nums[i] < nums[right] {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}

	nums[left], nums[right] = nums[right], nums[left]

	sortColors(nums[:left])
	sortColors(nums[left+1:])
}
