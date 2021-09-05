package main

import (
	"fmt"
	"time"
)

func main() {
	maps()
	arr()
}

func arr() bool {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("arr func %s\n", elapsed)
	}()
	count := 0
	value := "56254846579865465465498765465498"
	dp := make([]bool, 10)
	for _, v := range value {

		if !dp[v-'0'] {
			dp[v-'0'] = true
			count++
			if count >= 3 {
				return true
			}
		}
	}
	return false
}

func maps() bool {
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		fmt.Printf("maps func %s\n", elapsed)
	}()
	count := 0
	value := "56254846579865465465498765465498"
	dp := make(map[rune]bool, 10)
	for _, v := range value {

		if !dp[v-'0'] {
			dp[v-'0'] = true
			count++
			if count >= 3 {
				return true
			}
		}
	}
	return false
}
