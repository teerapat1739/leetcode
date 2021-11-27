package leetcode


func checkPossibility(nums []int) bool {
	// how to increase
	count := 0
	temp := nums[0]
	for i := 0; i < len(nums) - 1; i++ {
		if temp > nums[i + 1] {
			count++
		}else {
			temp = nums[i+1]
		}
		if count > 1 {
			return false
		}
	}
	return true
}

/*
3,4,2,3,
3,4,4,



*/
