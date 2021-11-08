package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximum Product of Two Elements in an Array.
Memory Usage: 3 MB, less than 100.00% of Go online submissions for Maximum Product of Two Elements in an Array.
*/

func maxProduct1464(nums []int) int {
	twoMaxVal := make([]int, 2)

	if nums[0] > nums[1] {
		twoMaxVal[0] = nums[1]
		twoMaxVal[1] = nums[0]
	} else {
		twoMaxVal[0] = nums[0]
		twoMaxVal[1] = nums[1]
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > twoMaxVal[0] && nums[i] > twoMaxVal[1] {
			twoMaxVal[0] = twoMaxVal[1]
			twoMaxVal[1] = nums[i]
		} else if nums[i] > twoMaxVal[0] && nums[i] <= twoMaxVal[1] {
			twoMaxVal[0] = nums[i]
		}
	}

	return (twoMaxVal[0] - 1) * (twoMaxVal[1] - 1)
}
