package leetcode

/*
Runtime: 144 ms, faster than 90.53% of Go online submissions for Monotonic Array.
Memory Usage: 9.5 MB, less than 85.26% of Go online submissions for Monotonic Array.

*/
func isMonotonic(nums []int) bool {
	isIncreasing, isDecreasing := true, true

	for i := 1; i < len(nums); i++ {
		if nums[i-1] > nums[i] {
			isIncreasing = false
		}

		if nums[i-1] < nums[i] {
			isDecreasing = false
		}
	}
	return isIncreasing || isDecreasing
}
