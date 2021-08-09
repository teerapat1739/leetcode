package leetcode

import (
	"strconv"
)

func syracuse(n int) string {
	// n is odd
	str := (strconv.Itoa(n))
	for {
		if n%2 != 0 {
			n = n*3 + 1
		} else {
			n = n / 2
		}

		if n == 1 {
			str = str + (" " + strconv.Itoa(n))
			break
		}
		str = str + (" " + strconv.Itoa(n))
	}
	return str
}
