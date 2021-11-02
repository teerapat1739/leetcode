package leetcode

/*
Runtime: 8 ms, faster than 87.77% of Go online submissions for Concatenation of Array.
Memory Usage: 6.4 MB, less than 76.36% of Go online submissions for Concatenation of Array.
*/
func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)
	return nums
}
