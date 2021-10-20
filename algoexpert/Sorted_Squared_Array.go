package algoexpert

func SortedSquaredArray(array []int) []int {
	// Write your code here.
	start := 0
	end := len(array) - 1
	dp := make([]int, len(array))

	for i := len(array) - 1; i >= 0; i-- {

		smallerVal := array[start]
		largerVal := array[end]
		if abs(smallerVal) > abs(largerVal) {
			dp[i] = smallerVal * smallerVal
			start++
		} else {
			dp[i] = largerVal * largerVal
			end--
		}
	}
	return dp
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
