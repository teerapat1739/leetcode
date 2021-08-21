package leetcode

func customSortString(order string, s string) string {
	orderMap := make(map[rune]int)
	sMap := make(map[rune]int)
	for _, v := range order {
		orderMap[v] += 1
	}

	for _, v := range s {
		sMap[v] += 1
	}
	res := []rune{}
	for _, v := range order {
		idxx := rune(v)
		_, Sok := sMap[idxx]
		if Sok {
			for sMap[idxx] != 0 {
				res = append(res, idxx)
				sMap[idxx] -= 1

			}
			delete(sMap, idxx)
		}
	}
	for i, _ := range sMap {
		for sMap[i] != 0 {
			res = append(res, rune(i))
			sMap[i] -= 1
			if sMap[i] == 0 {
				delete(sMap, i)
			}
		}
	}
	return string(res)
}
