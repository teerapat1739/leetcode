package leetcode

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	right := 0
	sum := 0
	ans := math.MaxInt64
	for right = 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target && left <= right {
			ans = int(math.Min(float64(ans), float64(right-left+1)))
			sum -= nums[left]
			left++
		}
	}
	if ans > len(nums) {
		return 0
	}
	return ans
}
