package leetcode

/*
Runtime: 108 ms, faster than 6.82% of Go online submissions for Maximum Subarray.
Memory Usage: 9.9 MB, less than 15.78% of Go online submissions for Maximum Subarray.
*/
// use sliding window technique
func maxSubArray(nums []int) int {
	maxSub := nums[0]
	currTotal := 0
	for _, num := range nums {
		if currTotal < 0 {
			currTotal = 0
		}
		currTotal += num
		maxSub = Max(currTotal, maxSub)

	}
	return maxSub
}
