package leetcode

/*
Success
Details
Runtime: 12 ms, faster than 90.20% of Go online submissions for Jump Game II.
Memory Usage: 5.6 MB, less than 83.75% of Go online submissions for Jump Game II.
Next challenges:

*/
func jump(nums []int) int {
	count := 0
	start, end := 0, 0

	for end < len(nums)-1 {
		nextPos := 0
		for i := start; i < end+1; i++ {
			nextPos = Max(nextPos, i+nums[i])
		}
		start = end + 1
		end = nextPos
		count++
	}
	return count
}

// 1,1,2,2,0,1,1

/*
2 >= 7
4 >= 7
  [4] == 0
  [3] --> 3+2
[5] >=7


*/

// 1,2,,0
