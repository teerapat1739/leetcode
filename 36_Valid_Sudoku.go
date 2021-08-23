package leetcode

func isValidSudoku(board [][]byte) bool {
	boardInt := [9][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				boardInt[i][j] = int(board[i][j] - '0')
			} else {
				boardInt[i][j] = 0
			}

		}
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			rowPass := checkRow(boardInt, i, j)
			colPass := checkCol(boardInt, i, j)
			boxPass := checkBox(boardInt, i, j)
			isNonZero := boardInt[i][j] != 0
			if isNonZero && (!rowPass || !colPass || !boxPass) {
				return false
			}
		}
	}
	return true
}

func checkRow(board [9][9]int, i, j int) bool {
	for k := 0; k < 9; k++ {
		isPass := board[i][k] == board[i][j]
		isEquelZero := board[i][k] == 0
		if !isEquelZero && isPass && k != j {
			return false
		}
	}
	return true
}

func checkCol(board [9][9]int, i, j int) bool {
	for k := 0; k < 9; k++ {
		isPass := board[k][j] == board[i][j]
		isEquelZero := board[k][j] == 0
		if !isEquelZero && isPass && k != i {
			return false
		}

	}
	return true
}

func checkBox(board [9][9]int, i, j int) bool {
	startRow := (i / 3) * 3
	startCol := (j / 3) * 3

	for k := startRow; k < startRow+3; k++ {
		for h := startCol; h < startCol+3; h++ {
			if board[k][h] != 0 && board[k][h] == board[i][j] && i != k && j != h {
				return false
			}
		}
	}
	return true
}
