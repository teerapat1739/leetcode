package leetcode

/*
Runtime: 8 ms, faster than 82.12% of Go online submissions for Peak Index in a Mountain Array.
Memory Usage: 4.6 MB, less than 39.11% of Go online submissions for Peak Index in a Mountain Array.
*/
func peakIndexInMountainArray(arr []int) int {

	i := 0
	for i = 1; i < len(arr)-1; i++ {
		for arr[i-1] < arr[i] {
			i++
		}
		if arr[i-1] > arr[i] {
			i--
			break
		}
	}
	return i
}
