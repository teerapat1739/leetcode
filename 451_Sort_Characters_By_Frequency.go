package leetcode

import (
	"sort"
	"strings"
)

type Row struct {
	c byte
	n int
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sort Characters By Frequency.
Memory Usage: 7.7 MB, less than 17.35% of Go online submissions for Sort Characters By Frequency.
*/

func frequencySort(s string) string {
	f := make([]int, 128)
	for i := 0; i < len(s); i++ {
		f[s[i]]++
	}
	var table []Row
	var res string
	for i, v := range f {
		if v != 0 {
			table = append(table, Row{byte(i), v})
		}
	}
	sort.Slice(table, func(i, j int) bool {
		return table[i].n > table[j].n
	})
	for _, v := range table {
		res += strings.Repeat(string(v.c), v.n)
		// n := v.n
		// for n != 0 {
		//     res += string(v.c)
		//     n--
		// }
	}
	return res
}
