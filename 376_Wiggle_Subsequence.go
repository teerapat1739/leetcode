package leetcode

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	count := 2
	previous := nums[1] - nums[0]
	if previous == 0 {
		count = 1
	}
	for i := 2; i < len(nums); i++ {
		current := nums[i] - nums[i-1]
		if (current > 0 && previous <= 0) || (current < 0 && previous >= 0) {
			count++
			previous = current
		}
	}
	return count
}
