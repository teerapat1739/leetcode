package leetcode

// time limit exceeded
func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs2(n int) int {
	memo := make([]int, n+1)
	return sum(n, memo)
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
Memory Usage: 2.1 MB, less than 21.35% of Go online submissions for Climbing Stairs.
Next challenges:
*/

func sum(n int, memo []int) int {
	if n == 0 || n == 1 {
		memo[n] = 1
		return 1
	}

	if memo[n] > 0 {
		return memo[n]
	}
	curr := sum(n-1, memo) + sum(n-2, memo)
	memo[n] = curr
	return curr
}

//4

/*
1+1+1+1
1+1+2
1+2+1
2+1+1
2+2
*/
