package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	dictS := make([]int, 26)
	for _, val := range s {
		pos := val - 97
		dictS[pos]++
	}

	dictT := make([]int, 26)
	for _, val := range t {
		pos := val - 97
		dictT[pos]++
	}

	for i, _ := range dictS {
		if dictS[i] != dictT[i] {
			return false
		}
	}

	return true
}
