package leetcode

/*

Constraints:
1 <= nums.length <= 2 * 104**4
1 <= nums[i] <= 104
Runtime: 4 ms, faster than 93.21% of Go online submissions for Delete and Earn.
Memory Usage: 3.7 MB, less than 75.31% of Go online submissions for Delete and Earn.
*/
func deleteAndEarn(nums []int) int {
	// make array size 2 * 104**4
	//capacity := 2 * math.Pow(float64(104), float64(4))
	dp1 := make(map[int]int)
	//dp1 := make([]int, 9)
	maxVal := 0

	for _, num := range nums {
		dp1[num] += num
		maxVal = max(maxVal, num)
	}
	dp2 := make([]int, maxVal+1)
	dp2[0] = dp1[0]
	dp2[1] = dp1[1]

	for i := 2; i < len(dp2); i++ {
		dp2[i] = max(dp2[i-1], dp2[i-2]+dp1[i])
	}

	return dp2[maxVal]
}

// 740. Delete and Earn
/*
[3,4,2,5,6,3]
2,3,3,4,5,6

delete 6 to earn 6 points, and also delete 5 : arr = [3,4,2,3]
delete 3 to earn 3 points and also delete 2,4: arr = [3]
delete 3 to earn 3 points and also delete 2,4: arr = []




*/
