package leetcode

import "fmt"

func optimalDivision(nums []int) string {

	if len(nums) == 1 {
		return fmt.Sprintf("%d", nums[0])
	}

	if len(nums) == 2 {
		return fmt.Sprintf("%d/%d", nums[0], nums[1])
	}

	result := fmt.Sprintf("%d/(", nums[0])

	for i := 1; i < len(nums); i++ {
		if i == len(nums)-1 {
			result += fmt.Sprintf("%d)", nums[i])
		} else {
			result += fmt.Sprintf("%d/", nums[i])
		}

	}

	return result
}
