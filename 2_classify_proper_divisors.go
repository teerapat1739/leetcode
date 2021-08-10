package leetcode

func classifyProperdisvors(number int) string {
	haft := number / 2
	divisors := 0
	for i := 1; i <= haft; i++ {
		if number%i == 0 {
			divisors += i
		}
		if divisors == number {
			return "perfect"
		}
		if divisors > number {
			return "abundant"
		}
	}
	return "deficient"
}
