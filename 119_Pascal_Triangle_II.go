package leetcode

/*
Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Pascal's Triangle II.
Memory Usage: 2.4 MB, less than 6.67% of Go online submissions for Pascal's Triangle II.
*/
func getRow(rowIndex int) []int {

	pascalArr := generate(rowIndex + 1)

	return pascalArr[rowIndex]
}
