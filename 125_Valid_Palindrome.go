package leetcode

import (
	"fmt"
	"strings"
)

/*
Success
Details
Runtime: 4 ms, faster than 67.23% of Go online submissions for Valid Palindrome.
Memory Usage: 3.3 MB, less than 48.13% of Go online submissions for Valid Palindrome.
*/
func isPalindrome(s string) bool {
	dp := make([]rune, 0, len(s))
	s = strings.ToLower(s)
	for _, c := range s {
		fmt.Println(c)
		if isvalidCharactter(c) {
			dp = append(dp, c)
		}
	}
	for i := 0; i < len(dp)/2; i++ {
		if dp[i] != dp[len(dp)-1-i] {
			return false
		}
	}
	return true
}

func isvalidCharactter(s rune) bool {
	return s >= '0' && s <= '9' || s >= 'a' && s <= 'z' || s >= 'A' && s <= 'Z'
}
