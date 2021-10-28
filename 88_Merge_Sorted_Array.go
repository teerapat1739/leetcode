package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Merge Sorted Array.
Memory Usage: 2.5 MB, less than 47.17% of Go online submissions for Merge Sorted Array.
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums3 := make([]int, len(nums1))
	copy(nums3, nums1)
	nums3 = nums3[:m]
	nums2 = nums2[:n]
	i, j := 0, 0
	ans := []int{}
	for i < len(nums3) || j < len(nums2) {
		if i >= len(nums3) {
			ans = append(ans, nums2[j])
			j++
			continue
		}
		if j >= len(nums2) {
			ans = append(ans, nums3[i])
			i++
			continue
		}

		if nums3[i] < nums2[j] {
			ans = append(ans, nums3[i])
			i++
		} else {
			ans = append(ans, nums2[j])
			j++
		}
	}
	for i := 0; i < len(nums1) && i < len(ans); i++ {
		nums1[i] = ans[i]
	}
}
