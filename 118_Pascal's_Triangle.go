package leetcode

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle.
Memory Usage: 2.3 MB, less than 7.63% of Go online submissions for Pascal's Triangle.
Next challenges:
*/
func generate(numRows int) [][]int {
	res := [][]int{
		[]int{1},
	}
	for i := 0; i < numRows-1; i++ {
		arrlen := len(res)
		temp := []int{0}
		temp = append(temp, res[arrlen-1]...)
		temp = append(temp, 0)
		row := []int{}
		for j := 0; j < len(res[arrlen-1])+1; j++ {
			row = append(row, temp[j]+temp[j+1])
		}
		res = append(res, row)
	}

	return res
}
