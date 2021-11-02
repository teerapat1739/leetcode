package leetcode

func buildArray1920(nums []int) []int {
	ans := make([]int, len(nums))

	for i, _ := range nums {
		ans[i] = nums[nums[i]]
	}
	return ans
}
