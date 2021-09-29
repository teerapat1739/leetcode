package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Is Subsequence.
Memory Usage: 2 MB, less than 51.98% of Go online submissions for Is Subsequence.
*/
func isSubsequence(s string, t string) bool {
	if len(t) < len(s) {
		return false
	}

	if len(s) == 0 {
		return true
	}
	current := 0
	for _, v := range t {
		currStr := s[current]
		if string(v) == string(currStr) {
			current++
		}

		if current >= len(s) {
			return true
		}
	}

	if current == len(s) {
		return true
	}

	return false
}
