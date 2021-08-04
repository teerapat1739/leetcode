package leetcode

import (
	"fmt"
	"math"
)

func minFallingPathSum(matrix [][]int) int {
	type Node struct {
		value    int
		position int
	}
	var n []Node
	if len(matrix) == 0 {
		return 0
	}
	sum := 0
	for i := 0; i < len(matrix); i++ {
		min := matrix[i][0]
		pos := 0
		pos2 := 0
		if len(n) > 0 && i > 1 {
			pos2 = n[len(n)-1].position
			matrix[i][pos] = 101
		}
		fmt.Println(matrix, pos2)
		for j := 0; j < len(matrix); j++ {
			fmt.Println("ddd", i, j, pos2, min, matrix[i][j])

			if (math.Abs(float64(j-pos2)) == 1) && i > 0 && len(matrix) > 2 {
				if j+1 > len(matrix) {
					continue
				}
				min = matrix[i][j]
				fmt.Println("mmm", min)
				continue
			}
			if matrix[i][j] < min {
				min = matrix[i][j]
				pos = j
			}

		}
		n = append(n, Node{min, pos})
	}
	fmt.Println("last", n)
	for i := 0; i < len(n); i++ {
		sum += n[i].value
	}
	return sum
}
