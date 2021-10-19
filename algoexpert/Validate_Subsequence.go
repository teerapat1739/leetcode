package algoexpert

func IsValidSubsequence(array []int, sequence []int) bool {
	// Write your code here.
	start := 0
	dp := make([]bool, len(sequence))
	for i := 0; i < len(array); i++ {
		if array[i] == sequence[start] {
			dp[start] = true
			start++
		}
	}
	return dp[len(sequence)-1]
}
