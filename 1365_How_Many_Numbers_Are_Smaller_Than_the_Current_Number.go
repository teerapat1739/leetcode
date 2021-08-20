package leetcode

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		cur_val := nums[i]
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < cur_val {
				count += 1
			}
		}
		result = append(result, count)
	}
	return result
}
