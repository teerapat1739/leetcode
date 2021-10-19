package algoexpert

import "sort"

func TwoNumberSum(array []int, target int) []int {

	// Write your code here.
	sort.Ints(array)

	left, right := 0, len(array)-1
	var ans []int
	for left < right {
		sum := array[left] + array[right]
		if sum == target {
			ans = []int{
				array[left], array[right],
			}
			break
		}
		if sum < target {
			left++
			continue
		}

		if sum > target {
			right--
			continue
		}
	}
	return ans
}
