package leetcode

func validSameNumber(a, b int) string {

	for {

		modA := a % 10
		modB := b % 10
		if modA == modB {
			return "Invalid"
		}

		a = a / 10
		b = b / 10
		if a < 10 || b < 10 {
			if a == b {
				return "Invalid"
			}
			break
		}
	}

	return "Valid"
}
