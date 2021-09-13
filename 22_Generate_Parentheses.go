package leetcode

import (
	"strings"
)

var stack, res []string
func generateParenthesis(n int) []string {
	 stack, res = []string{},[]string{}
	// only add open parenthesis if open < n
	// only add a closing parenthesis if closed < open
	// valid: open == closed == n

	return backtrack(0, 0, n)
}

func backtrack(openN, closeN, n int) []string {
	// valid: open == closed == n
	if openN == closeN && openN == n {
		res = append(res, strings.Join(stack,""))
		//fmt.Println(res)
		return res
	}
	// only add open parenthesis if open < n
	if openN < n {
		stack = append(stack, "(")
		backtrack(openN+1, closeN, n)
		//pop
		stack = stack[:len(stack)-1]
	}

	// only add a closing parenthesis if closed < open
	if closeN < openN {
		stack = append(stack, ")")
		backtrack(openN, closeN+1, n)
		//pop
		stack = stack[:len(stack)-1]
	}

	return res
}
