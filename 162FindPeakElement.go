package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Find Peak Element.
Memory Usage: 2.8 MB, less than 100.00% of Go online submissions for Find Peak Element.
*/
func findPeakElement(nums []int) int {
	peakIdx := 0
	i := 0
	for i = 0; i < len(nums)-1; i++ {
		for i < len(nums)-1 && nums[i] < nums[i+1] {
			i++
		}
		peakIdx = i
		for i < len(nums)-1 && nums[i] > nums[i+1] {
			i++
		}
	}
	return peakIdx
}
