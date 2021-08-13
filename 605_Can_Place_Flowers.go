package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 0 {
		return false
	}

	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			return true
		}

	}

	m := make(map[int]int)

	for i := 0; i < len(flowerbed); i++ {
		m[i] = flowerbed[i]
	}

	temp := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			temp++
		} else {
			temp = 0
		}

		if temp == 1 || temp == 2 {
			v1, ok1 := m[i-1]
			v2, ok2 := m[i+1]
			if (ok2 && v2 == 0) && (!ok1 && v1 == 0) {
				m[i] = 1
				n--
				temp = 0
			} else if i+1-temp == 0 || i+1 == len(flowerbed) {
				m[i] = 1
				n--
				temp = 0
			} else if (ok2 && v2 == 0) && (ok1 && v1 == 0) {
				m[i] = 1
				n--
				temp = 0
			}
		}
	}
	return n <= 0
}
