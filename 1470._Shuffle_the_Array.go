package leetcode

/*
Runtime: 4 ms, faster than 88.48% of Go online submissions for Shuffle the Array.
Memory Usage: 3.9 MB, less than 35.15% of Go online submissions for Shuffle the Array.
*/
func shuffle(nums []int, n int) []int {
	ans := []int{}
	for i := 0; i < n; i++ {
		ans = append(ans, []int{
			nums[i],
			nums[i+n],
		}...)
	}
	return ans
}
