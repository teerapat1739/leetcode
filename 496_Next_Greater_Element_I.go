package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	dp := make(map[int]int)
	var stack []int
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			popValue := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			dp[popValue] = nums2[i]
		}
		stack = append(stack, nums2[i])
	}

	ans := make([]int, len(nums1))
	for idx, val := range nums1 {
		v, ok := dp[val]
		if !ok {
			ans[idx] = -1
		} else {
			ans[idx] = v
		}
	}
	return ans
}
