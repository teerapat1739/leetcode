package leetcode

import "fmt"

func nextGreaterElements(nums []int) []int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
	return []int{}
}
