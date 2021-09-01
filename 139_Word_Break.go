package leetcode

// Problem use dynamic programming(iteration) to solve the problem
/*
Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Word Break.
Memory Usage: 2.2 MB, less than 55.69% of Go online submissions for Word Break.
*/
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	mapWordDict := make(map[string]bool)
	for _, word := range wordDict {
		mapWordDict[word] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			substr := s[j:i]
			if dp[j] && mapWordDict[substr] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(dp)-1]
}
