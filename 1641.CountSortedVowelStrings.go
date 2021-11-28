package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Count Sorted Vowel Strings.
Memory Usage: 1.9 MB, less than 75.61% of Go online submissions for Count Sorted Vowel Strings.
*/
func countVowelStrings(n int) int {
	if n == 1 {
		return 5
	}

	dp := [5]int{1, 2, 3, 4, 5}
	for n > 1 {
		for i := 0; i < 5; i++ {
			dp[i] += dp[i-1]
		}
		n--
	}

	var ans int

	for i := 0; i < 5; i++ {
		ans += dp[i]
	}
	return ans
}
