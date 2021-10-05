package leetcode

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	count := 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			count += i
			n /= i
		}
	}
	return count

}
