package leetcode

func stoneGame(piles []int) bool {
	len := len(piles)

	start, end := 0, len-1
	for {
		if start < end {
			break
		}

	}
	return true
}

/*
5 3 4 5
Alice: 5 4
BoB  : 5 3

3,7,2,3,
A: 3 7
B: 3 2

7,8,8,10,
A: 10
B:


5, 33, 11, 4, 5, 7, 5 , 3, 7, 11 ,9, 17
A: 11 11 5
B: 9  7  33
*/

//[3 7,8, 6, 4, 8,10 3]
/*
10 7 6
8 8  4
*/

// 	aliceArr := []int{}
// 	bobArr := []int{}

// 	left := 0
// 	right := len(piles) - 1
// 	for len(piles) > len(aliceArr) + len(bobArr) {
// 		val := 0
// 		if piles[left] == piles[right] {
// 			tempR := right - 1
// 			tempL := left + 1
// 			for piles[tempR] == piles[tempL] {
// 				tempR = right - 1
// 				tempL = left + 1
// 			}
// 			if piles[tempR] > piles[tempL] {
// 				val = piles[left]
// 				left++
// 			}else {
// 				val = piles[right]
// 				right--
// 			}
// 		}else if piles[left] > piles[right] {
// 			val = piles[left]
// 			left++
// 		}else {
// 			val = piles[right]
// 			right--
// 		}

// 		if len(bobArr) < len(aliceArr) {
// 			bobArr = append(bobArr, val)
// 		}else {
// 			aliceArr = append(aliceArr, val)
// 		}

// 	}

// 	sumAlice := 0
// 	sumBob := 0
// 	for i := 0; i < len(bobArr); i++ {
// 		sumAlice += aliceArr[i]
// 		sumBob += bobArr[i]
// 	}
// 	return sumAlice > sumBob
