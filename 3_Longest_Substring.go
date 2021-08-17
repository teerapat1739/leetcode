package leetcode

func lengthOfLongestSubstring(s string) int {
	m := make(map[string]int)
	count := 0
	max := 0
	str := ""
	for i := 0; i < len(s); i++ {
		v, ok := m[string(s[i])]
		if ok {
			m = make(map[string]int)
			i = v
			str = ""
			count = 0
		} else {
			m[string(s[i])] = i
			str += string(s[i])
			count++
		}
		if count > max {
			max = count
		}
	}
	return max
}
