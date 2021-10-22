package algoexpert

func TournamentWinner(competitions [][]string, results []int) string {
	// Write your code here.
	dp := make(map[string]int)

	dict := make(map[int]int, 2)
	dict[0] = 1
	dict[1] = 0

	for i := 0; i < len(competitions); i++ {
		v := dict[results[i]]
		teamWin := competitions[i][v]
		dp[teamWin] += 3
	}

	teamWin := ""
	pointWin := 0
	for key, element := range dp {
		if element > pointWin {
			pointWin = element
			teamWin = key
		}
	}
	return teamWin
}
