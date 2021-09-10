package leetcode

/*
Success
Details
Runtime: 4 ms, faster than 81.75% of Go online submissions for Search Insert Position.
Memory Usage: 3 MB, less than 28.19% of Go online submissions for Search Insert Position.

*/
func searchInsert(nums []int, target int) int {
	count := 0
	for i, v := range nums {
		if target <= v {
			return i
		}
		if i == len(nums)-1 {
			return len(nums)
		}
	}
	return count
}
