package leetcode

func maxScoreSightseeingPair2(values []int) int {
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

func maxScoreSightseeingPair(values []int) int {
	i := 0
	ans := 0
	for j := 1; j < len(values); j++ {
		ans = max(ans, values[i]+values[j]+i-j)
		if values[i]+values[j]+i-j < 0 {
			i = j
		}
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
