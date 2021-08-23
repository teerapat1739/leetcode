package leetcode

func areOccurrencesEqual(s string) bool {

	if len(s) == 0 {
		return false
	}

	m := make(map[rune]int)
	for _, v := range s {
		m[v] += 1
	}

	start := m[rune(s[0])]

	for _, v := range m {
		if v != start {
			return false
		}
	}
	return true
}
