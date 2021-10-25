package leetcode

import "strings"

/*
Runtime: 691 ms, faster than 19.05% of Go online submissions for License Key Formatting.
Memory Usage: 40.2 MB, less than 7.14% of Go online submissions for License Key Formatting.
*/

func licenseKeyFormatting(s string, k int) string {
	ans := ""
	count := 0
	s = strings.ReplaceAll(s, "-", "")
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "-" {
			continue
		}

		if count%k == 0 && i != len(s)-1 {
			ans = "-" + ans
		}
		ans = strings.ToUpper(string(s[i])) + ans
		count++
	}

	return ans
}

/*
Input: s = "5F3Z-2e-9-w", k = 4
Output: "5F3Z-2E9W"
Explanation: The string s has been split into two parts, each part has 4 characters.
Note that the two extra dashes are not needed and can be removed.
*/
