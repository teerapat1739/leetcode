package leetcode

func maxScoreSightseeingPair(values []int) int {
	res := 0
	pos := 0
	for i := 1; i < len(values); i++ {
		res = max(res, values[i]+values[pos]+pos-i)
		if values[i]+i-pos > values[pos] {
			pos = i
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
