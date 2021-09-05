package leetcode

//	Time Limit Exceeded
func longestPalindrome(s string) string {
	sub := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isPalindrome(s[i : j+1]) {
				if len(sub) < len(s[i:j+1]) {
					sub = s[i : j+1]
				}
			}
		}
	}
	return sub
}
