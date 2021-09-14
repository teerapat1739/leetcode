package leetcode

/*
Success
Details
Runtime: 56 ms, faster than 93.18% of Go online submissions for Jump Game.
Memory Usage: 7.1 MB, less than 86.79% of Go online submissions for Jump Game.


*/
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	count := 0
	for i := len(nums) - 2; i >= 0; i-- {

		if nums[i] > 0 {
			if count == 0 {
				continue
			}
			if nums[i] >= count+1 {
				count = 0
			} else {
				count++
			}

		} else {
			count++
		}

	}

	return count == 0
}
