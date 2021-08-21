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
	for _, v := range orderMap {
		idxx := rune(v)
		iSval, Sok := sMap[idxx]
		if Sok {
			sMap[idxx] -= 1
			if iSval == 0 {
				delete(sMap, idxx)
			}
			res = append(res, idxx)
		}
	}

	for _, v := range sMap {
		idxx := rune(v)
		_, Oderok := orderMap[idxx]
		if Oderok {
			sMap[idxx] -= 1
			if sMap[idxx] == 0 {
				delete(sMap, idxx)
			}
			res = append(res, idxx)
		}
	}
	return string(res)
}
