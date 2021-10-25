package leetcode

import "sort"

/*
Runtime: 24 ms, faster than 87.71% of Go online submissions for Group Anagrams.
Memory Usage: 8.6 MB, less than 46.72% of Go online submissions for Group Anagrams.
*/
func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)

	for _, str := range strs {
		b := []rune(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		keys := string(b)
		dict[keys] = append(dict[keys], str)
	}

	var ans [][]string

	for _, val := range dict {
		ans = append(ans, val)
	}

	return ans
}

/*
Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/
