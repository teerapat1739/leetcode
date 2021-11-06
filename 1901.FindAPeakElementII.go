package leetcode

/*
Runtime: 164 ms, faster than 100.00% of Go online submissions for Find a Peak Element II.
Memory Usage: 11.9 MB, less than 84.00% of Go online submissions for Find a Peak Element II.
*/
func findPeakGrid(mat [][]int) []int {
	for row := 0; row < len(mat); row++ {
		for col := 0; col < len(mat[row]); col++ {
			top := getVal(row-1, col, mat)
			bottom := getVal(row+1, col, mat)
			right := getVal(row, col+1, mat)
			left := getVal(row, col-1, mat)
			if compareX(mat[row][col], top, right, bottom, left) {
				return []int{
					row,
					col,
				}
			}
		}
	}
	return []int{}

}

func compareX(target int, lists ...int) bool {
	for _, list := range lists {
		if target <= list {
			return false
		}
	}
	return true
}

func getVal(row, col int, mat [][]int) int {
	if col < 0 || row < 0 || row >= len(mat) || col >= len(mat[0]) {
		return -1
	}

	return mat[row][col]
}
