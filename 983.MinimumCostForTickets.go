package leetcode


func mincostTickets(days []int, costs []int) int {
	n := len(days)
	dp := make([]int, n + 1)
	//for i, _ := range dp {
	//	dp[i] = costs[0]*365
	//}
	for i := n - 1; i >=0; i-- {
		d7 := i
		d30 := i
		dd7 := days[d7]
		dPL7 := days[i] + 7
		for d7 < n && dd7 < dPL7 {
			d7++
		}
		dd30 := days[d30]
		dPL30 := days[i] + 30
		for d30 < n && dd30 < dPL30 {
			d30++
		}
		costsD1 := costs[0] + dp[i+1]
		costsD7 := costs[1] + dp[d7]
		costsD30 := costs[2] + dp[d30]

		dp[i] = min(costsD1, min(costsD7, costsD30))
	}

	return dp[0]
}