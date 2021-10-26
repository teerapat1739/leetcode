package leetcode

import "sort"

// Time Limit Exceeded
func findAnagrams438(s string, p string) []int {
	a := []rune(p)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	var ans []int
	for i := 0; i <= len(s)-len(p); i++ {
		str := s[i : i+len(p)]
		b := []rune(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		if string(a) == string(b) {
			ans = append(ans, i)
		}
	}
	return ans
}

func findAnagrams438II(s string, p string) []int {
	var ans []int
	dict := make([]int, 26)
	for _, val := range p {
		dict[val-97]++
	}
	lenP := len(p)
	tempLenP := lenP
	var tempDict []int
	tempDict = append(tempDict, dict...)
	i := 0
	lenS := len(s)
	for i < lenS {
		if tempDict[s[i]-97] == 0 {
			tempDict = []int{}
			tempDict = append(tempDict, dict...)
			if tempLenP == lenP {
				i++
			}
			tempLenP = lenP

			continue
		}
		tempLenP--
		tempDict[s[i]-97]--
		if tempLenP == 0 {
			start := i - lenP + 1
			i = start
			ans = append(ans, start)
			tempDict = []int{}
			tempDict = append(tempDict, dict...)
			tempLenP = lenP
		}
		i++
	}

	return ans
}
