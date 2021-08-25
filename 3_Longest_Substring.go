package leetcode

import (
	"math"
)

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

// solve with sliding window algo

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	max := 1
	m := make(map[string]bool)
	left := 0
	right := 0
	for right < len(s) {
		for left < right && m[string(s[right])] {
			m[string(s[left])] = false
			left++
		}
		m[string(s[right])] = true
		max = int(math.Max(float64(max), float64(right-left+1)))
		right++
	}

	return max
}
