package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Spiral Matrix.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Spiral Matrix.
*/
func spiralOrder(matrix [][]int) []int {
	row, col := len(matrix), len(matrix[0])

	tempStartRow, tempEndRow := 0, row
	tempStartCol, tempEndCol := 0, col

	curRow, curCol := 0, 0

	ans := []int{}
	for {

		for curCol < tempEndCol && len(ans) < row*col {
			ans = append(ans, matrix[curRow][curCol])
			curCol++
		}
		curCol--
		curRow++
		for curRow < tempEndRow && len(ans) < row*col {
			ans = append(ans, matrix[curRow][curCol])
			curRow++
		}
		curRow--
		curCol--
		for curCol >= tempStartCol && len(ans) < row*col {
			ans = append(ans, matrix[curRow][curCol])
			curCol--
		}
		curCol++
		curRow--
		for curRow > tempStartRow && len(ans) < row*col {
			ans = append(ans, matrix[curRow][curCol])
			curRow--
		}
		curRow++
		curCol++

		if len(ans) >= row*col {
			break
		}
		tempStartRow++
		tempStartCol++
		tempEndRow--
		tempEndCol--

	}
	return ans
}
