package leetcode

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

//4

/*
1+1+1+1
1+1+2
1+2+1
2+1+1
2+2
*/
