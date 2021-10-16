package leetcode

func maxProduct2(nums []int) int {
	currMax, currMin := 1, 1
	res := maxS(nums...)
	for _, num := range nums {
		if num == 0 {
			currMax, currMin = 1, 1
		}
		temp := currMax * num
		curMinNum := currMin * num
		currMax = maxS(currMax*num, num, curMinNum)
		currMin = minS(temp, curMinNum, num)
		res = max(res, currMax)
	}
	return res
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Maximum Product Subarray.
Memory Usage: 2.8 MB, less than 47.69% of Go online submissions for Maximum Product Subarray.
*/
func maxProduct(nums []int) int {
	maxOverall := nums[0]

	maxEndingHere, minEndingHere := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		temp := maxEndingHere
		maxEndingHere = maxS(nums[i], nums[i]*maxEndingHere, nums[i]*minEndingHere)
		minEndingHere = minS(nums[i], nums[i]*temp, nums[i]*minEndingHere)
		maxOverall = maxS(maxOverall, maxEndingHere)

	}
	return maxOverall
}

func maxS(n ...int) int {
	if len(n) == 0 {
		return 1
	}
	maxCurr := n[0]
	for _, i2 := range n {
		maxCurr = max(maxCurr, i2)
	}
	return maxCurr
}

func minS(n ...int) int {
	if len(n) == 0 {
		return 1
	}
	maxCurr := n[0]
	for _, i2 := range n {
		maxCurr = min(maxCurr, i2)
	}
	return maxCurr
}

/*
[2,3,-2,4]

dp[2]
dp[2,6]
-12 < 6
dp[2,6,-2,-8]


[-1,2,3,-4]

dp[-1]

dp[-1,2] (-2,2)

dp[-1,2,6] (-6,6)

dp[-1,2,6,24]



{-2,0,1}
dp[-2] (-2,-2)
dp[-2,0] (1,1)
dp[-2,0,-3]
each round
dp[-]



[0,2]

dp[0] (1,1)
dp[0,2]

[3,-1,4]
dp[3] (3,3)
dp[3, -1] (3,-3)
dp[]
*/

//func maxSubArray2(nums []int) int {
//	maxCurrent, maxGlobal := nums[0], nums[0]
//	for i := 1; i < len(nums); i++ {
//		maxCurrent = max(nums[i], maxCurrent+nums[i])
//		if maxCurrent > maxGlobal {
//			maxGlobal = maxCurrent
//		}
//	}
//	return maxGlobal
//}
