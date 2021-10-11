package leetcode

// https://leetcode.com/submissions/detail/550380378/
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

func longestPalindrome2(s string) string {

	sub := s[0:1]
	startP, endP := 0, 0
	end := len(s) - 1
	for i := startP; i < len(s); i++ {
		start := i
		for j := end; j > i; j-- {
			tempSub := ""
			end := j
			start = i
			startStr := string(s[start])
			endStr := string(s[end])
			for startStr == endStr && start < end {
				if end-start-1 <= 1 {
					tempSub = s[i : j+1]
					// startP = i
					endP = j
					break
				}
				if len(sub) == len(s) || endP+len(sub) >= len(s) {
					return sub
				}
				if len(sub) > end-start && startStr != endStr {
					break
				}
				start++
				end--
				startStr = string(s[start])
				endStr = string(s[end])
			}
			if len(tempSub) > 0 {
				sub = compare(tempSub, sub)
				i = start
				break
			}
		}
	}
	return sub
}

func compare(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}

func longestPalindrome3(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		odd := expand(s, i, i)
		even := expand(s, i, i+1)

		len := max(odd, even)
		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}
	return s[start : end+1]
}

func expand(s string, start, end int) int {
	for start >= 0 && end < len(s) && s[start] == s[end] {
		start--
		end++
	}
	return end - start - 1
}
