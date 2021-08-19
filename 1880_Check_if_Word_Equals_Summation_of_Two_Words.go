package leetcode

import (
	"math"
)

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	firstWordRune := []rune(firstWord)
	secondWordRune := []rune(secondWord)
	targetWordRune := []rune(targetWord)
	sumFirstWord := 0
	sumSecondWord := 0
	sumTargetWord := 0
	s := len(firstWord) - 1
	for _, v := range firstWordRune {
		mp := math.Pow(10, float64(s))
		sumFirstWord += int(math.Abs(float64((97 - int(v))) * mp))
		s--
	}
	s = len(secondWord) - 1
	for _, v := range secondWordRune {
		mp := math.Pow(10, float64(s))
		sumSecondWord += int(math.Abs(float64((97 - int(v))) * mp))
		s--
	}
	s = len(targetWord) - 1
	for _, v := range targetWordRune {
		mp := math.Pow(10, float64(s))
		sumTargetWord += int(math.Abs(float64((97 - int(v))) * mp))
		s--
	}
	return sumFirstWord+sumSecondWord == sumTargetWord
}
