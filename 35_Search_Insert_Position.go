package leetcode

/*
Success
Details
Runtime: 4 ms, faster than 81.75% of Go online submissions for Search Insert Position.
Memory Usage: 3 MB, less than 28.19% of Go online submissions for Search Insert Position.

*/
func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v == target || v > target {
			return i
		}
	}
	return len(nums)
}
