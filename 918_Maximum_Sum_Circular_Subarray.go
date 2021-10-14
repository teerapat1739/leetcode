package leetcode

import "math"

/*
Runtime: 60 ms, faster than 59.42% of Go online submissions for Maximum Sum Circular Subarray.
Memory Usage: 7.4 MB, less than 100.00% of Go online submissions for Maximum Sum Circular Subarray.
*/
func maxSubarraySumCircular(nums []int) int {
	maxN := math.MinInt16
	minN := math.MaxInt16
	curr1, curr2, total := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		curr1 += nums[i]
		maxN = max(maxN, curr1)
		curr1 = max(curr1, 0)

		curr2 += nums[i]
		minN = min(minN, curr2)
		curr2 = min(curr2, 0)

		total += nums[i]
	}

	if maxN < 0 {
		return maxN
	}
	return max(maxN, total-minN)
}

/*


 */
