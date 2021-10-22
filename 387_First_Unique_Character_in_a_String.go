package leetcode

/*
Runtime: 28 ms, faster than 63.01% of Go online submissions for First Unique Character in a String.
Memory Usage: 5.5 MB, less than 54.28% of Go online submissions for First Unique Character in a String.
*/
func firstUniqChar(s string) int {

	ans := -1
	mapQueue := make(map[rune]int)
	for i := 0; i < len(s); i++ {

		mapQueue[rune(s[i])] += 1
	}

	for i := 0; i < len(s); i++ {
		v := mapQueue[rune(s[i])]
		if v == 1 {
			return i
		}
	}

	return ans
}

/*
Runtime: 9 ms, faster than 77.14% of Go online submissions for First Unique Character in a String.
Memory Usage: 5.7 MB, less than 40.71% of Go online submissions for First Unique Character in a String.
*/
func firstUniqChar2(s string) int {
	counter := make([]int, 26)

	for _, r := range s {
		charIdx := r - 'a'
		counter[charIdx] += 1
	}

	for idx, r := range s {
		charIdx := r - 'a'
		count := counter[charIdx]
		if count == 1 {
			return idx
		}
	}

	return -1
}

/*
leetcode

l e e t c o d e
1 1 2 1 1 1 1 3

*/
