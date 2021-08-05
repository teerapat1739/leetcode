package leetcode

import (
	"math"
)

func minFallingPathSum(matrix [][]int) int {

	for i := 1; i < len(matrix); i++ {

		for j := 0; j < len(matrix[i]); j++ {
			arrPos := []int{j - 1, j, j + 1}
			min := 101
			for _, pos := range arrPos {
				if pos < 0 || pos >= len(matrix[i]) {
					continue
				}
				if matrix[i-1][pos] < min {
					min = matrix[i-1][pos]
				}
			}
			matrix[i][j] += min
		}
	}

	result := math.MaxInt64

	for _, v := range matrix[len(matrix)-1] {
		if v < result {
			result = v
		}
	}

	return result
}
