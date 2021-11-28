package leetcode

/*
Runtime: 33 ms, faster than 7.06% of Go online submissions for Minimum Falling Path Sum.
Memory Usage: 5.6 MB, less than 41.18% of Go online submissions for Minimum Falling Path Sum.
*/
func minFallingPathSum931(matrix [][]int) int {
	n := len(matrix)
	for row := 1; row < n; row++ {
		for col := 0; col < n; col++ {
			top := getVal931(row, col, matrix) + getVal931(row-1, col, matrix)
			diagonallyLeft := getVal931(row, col, matrix) + getVal931(row-1, col-1, matrix)
			diagonallyRight := getVal931(row, col, matrix) + getVal931(row-1, col+1, matrix)
			matrix[row][col] = min931(top, diagonallyLeft, diagonallyRight)
		}
	}
	return min931(matrix[n-1]...)
}

func getVal931(row, col int, matrix [][]int) int {
	if (row >= 0 && row < len(matrix)) && (col >= 0 && col < len(matrix)) {
		return matrix[row][col]
	}
	return 10000000
}

func min931(args ...int) int {
	ans := args[0]
	for _, v := range args {
		if ans > v {
			ans = v
		}
	}
	return ans
}
