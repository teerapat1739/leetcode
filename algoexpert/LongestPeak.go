package algoexpert

import "fmt"

func LongestPeak(array []int) int {
	// Write your code here.
	longestPeakLenght := 0
	i := 1
	for i < len(array)-1 {

		isCandidatePeak := array[i-1] < array[i] && array[i] > array[i+1]
		if !isCandidatePeak {
			i++
			continue
		}

		fmt.Println(i)

		leftIdx := i
		for leftIdx > 0 && array[leftIdx-1] < array[leftIdx] {
			leftIdx--
			longestPeakLenght := 0
			i := 1
			for i < len(array)-1 {

				isCandidatePeak := array[i-1] < array[i] && array[i] > array[i+1]
				if !isCandidatePeak {
					i++
					continue
				}

				fmt.Println(i)

				leftIdx := i
				for leftIdx > 0 && array[leftIdx-1] < array[leftIdx] {
					leftIdx--
				}

				rightIdx := i
				for rightIdx < len(array)-1 && array[rightIdx] > array[rightIdx+1] {
					rightIdx++
				}
				currentLongestPeakLenght := rightIdx - leftIdx + 1
				if currentLongestPeakLenght > longestPeakLenght {
					longestPeakLenght = currentLongestPeakLenght
				}
				i = rightIdx
			}

			return longestPeakLenght
		}

		rightIdx := i
		for rightIdx < len(array)-1 && array[rightIdx] > array[rightIdx+1] {
			rightIdx++
		}
		currentLongestPeakLenght := rightIdx - leftIdx + 1
		if currentLongestPeakLenght > longestPeakLenght {
			longestPeakLenght = currentLongestPeakLenght
		}
		i = rightIdx
	}

	return longestPeakLenght
}
