package leetcode

/*
Runtime: 40 ms, faster than 48.13% of Go online submissions for Rotate Array.
Memory Usage: 8.9 MB, less than 20.17% of Go online submissions for Rotate Array.
*/
func rotate(nums []int, k int) {
	k = k % len(nums)
	sizeNums := len(nums)
	temp := nums[sizeNums-k:]
	temp2 := nums[:sizeNums-k]
	temp = append(temp, temp2...)
	for i := 0; i < len(nums); i++ {
		nums[i] = temp[i]
	}

}
