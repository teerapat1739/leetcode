package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Build an Array With Stack Operations.
Memory Usage: 2.4 MB, less than 100.00% of Go online submissions for Build an Array With Stack Operations.
*/
func buildArray(target []int, n int) []string {
	var ans []string
	idx := 0
	for i := 0; i < n && idx < len(target); i++ {
		if i+1 == target[idx] {
			ans = append(ans, "Push")
			idx++
		} else {
			ans = append(ans, "Push")
			ans = append(ans, "Pop")
		}
	}
	return ans
}
