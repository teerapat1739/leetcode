package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Two Sum II - Input Array Is Sorted.
Memory Usage: 3 MB, less than 68.77% of Go online submissions for Two Sum II - Input Array Is Sorted.
*/

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

	return []int{left + 1, right + 1}
}
