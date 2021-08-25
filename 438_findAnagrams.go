package leetcode

// time limit exceeded
func findAnagrams(s string, p string) []int {
	if len(s) == 0 || len(p) == 0 || len(s) < len(p) {
		return []int{}
	}
	arr := []int{}
	lenP := len(p)
	start := 0
	end := lenP
	for end <= len(s) {
		// subStr := []byte(s[start:end])
		subStr := s[start:end]
		ispass := true
		stateMp := make(map[string]int)
		for _, v := range p {
			stateMp[string(v)] += 1
		}
		for _, v := range subStr {
			_, ok := stateMp[string(v)]
			if !ok {
				ispass = false
				break
			}
			stateMp[string(v)] -= 1
			if stateMp[string(v)] == 0 {
				delete(stateMp, string(v))
			}
		}

		if ispass {
			arr = append(arr, start)
		}
		start++
		end = start + lenP
	}
	return arr
}

/*
Success: Runtime: 720 ms, faster than 21.21% of Go online submissions for Find All Anagrams in a String.
Memory Usage: 5.2 MB, less than 55.76% of Go online submissions for Find All Anagrams in a String.
*/

func findAnagrams2(s string, p string) []int {
	// ms :=

	target := [26]int{}
	count := [26]int{}

	for _, v := range p {
		target[v-'a']++
	}
	arr := []int{}
	for i := 0; i < len(s); i++ {
		c := 0
		for c < len(p) && i+c < len(s) {
			count[s[i+c]-'a']++
			c++
		}
		if count == target {
			arr = append(arr, i)
		}
		count = [26]int{}
	}

	return arr
}

func findAnagrams3(s string, p string) []int {
	// ms :=
	if len(s) < len(p) {
		return []int{}
	}

	arr := []int{}
	target := [26]int{}
	count := [26]int{}

	for i := range p {
		target[p[i]-'a']++
		count[s[i]-'a']++
	}
	start := 0
	for i := len(p); i < len(s); i++ {
		if count == target {
			arr = append(arr, start)
		}
		count[s[start]-'a']--
		count[s[i]-'a']++
		start++
	}

	if count == target {
		arr = append(arr, start)
	}

	return arr
}
