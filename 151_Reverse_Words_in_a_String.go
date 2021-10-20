package leetcode

import (
	"fmt"
	"strings"
)

/*
Runtime: 6 ms, faster than 20.13% of Go online submissions for Reverse Words in a String.
Memory Usage: 8.5 MB, less than 8.80% of Go online submissions for Reverse Words in a String.
*/
func reverseWords(s string) string {
	s = strings.TrimSpace(fmt.Sprintf("\t %s\n ", s))
	var dp []string
	curr := ""
	for i := 0; i < len(s); i++ {
		if s[i] != 32 {
			curr += string(s[i])
		} else if len(curr) > 0 {
			dp = append(dp, curr)
			curr = ""
		}
	}
	if len(curr) > 0 {
		dp = append(dp, curr)
	}
	ans := ""
	for i := len(dp) - 1; i >= 0; i-- {
		ans += dp[i]
		if i > 0 {
			ans += " "
		}
	}

	return ans
}
