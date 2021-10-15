package leetcode

func maxProduct(nums []int) int {
	currMax, currMin := 1, 1
	res := maxS(nums...)
	for _, num := range nums {
		if num == 0 {
			currMax, currMin = 1, 1
		}
		temp := currMax * num
		currMax = maxS(currMax*num, num, currMin*num)
		currMin = minS(temp, num*currMin, num)
		res = max(res, currMax)
	}
	return res
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
