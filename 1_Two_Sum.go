package leetcode

import "sort"

/*
Runtime: 8 ms, faster than 64.74% of Go online submissions for Two Sum.
Memory Usage: 5.6 MB, less than 9.14% of Go online submissions for Two Sum.
*/

func twoSum2(nums []int, target int) []int {

	dict := make(map[int][]int, len(nums))

	for i, num := range nums {
		dict[num] = append(dict[num], i)
	}
	sort.Ints(nums)

	var ans []int
	left, right := 0, len(nums)-1
	for left < right {
		curr := nums[left] + nums[right]
		if curr == target {
			ansL, ansR := 0, 0
			ansL, dict[nums[left]] = Dequeue(dict[nums[left]])
			ansR, dict[nums[right]] = Dequeue(dict[nums[right]])
			ans = []int{
				ansL,
				ansR,
			}
			right--
			left++
		} else if curr > target {
			right--
		} else {
			left++
			// curr < target
		}
	}
	return ans
}

func Dequeue(arr []int) (int, []int) {
	if len(arr) == 0 {
		return 0, nil
	}
	num := arr[0]
	arr = arr[1:]
	return num, arr
}
