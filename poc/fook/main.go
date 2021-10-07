package main

import (
	"fmt"
	"math"
)

type Tx struct {
	Price float64
	Date  float64
}

type Interval struct {
	From float64
	To   float64
	tx   []Tx
}

func Fook() []Interval {

	txs := []Tx{
		{Price: 1, Date: 1},
		{Price: 3, Date: 2},
		{Price: 1.7, Date: 4},
	}

	intervals := []Interval{
		{From: 1, To: 2, tx: []Tx{}},
		{From: 2, To: 3, tx: []Tx{}},
		{From: 3, To: 4, tx: []Tx{}},
	}

	dpFrom := make(map[float64]int)
	dpTo := make(map[float64]int)

	for i, v := range intervals {
		dpFrom[v.From] = i
		dpTo[v.To] = i
	}

	for _, tx := range txs {
		valFrom, okFrom := dpFrom[tx.Date]
		valTo, okTo := dpTo[tx.Date]
		if okTo && okFrom {
			intervals[valFrom].tx = append(intervals[valFrom].tx, Tx{
				Price: tx.Price,
			})

			intervals[valTo].tx = append(intervals[valTo].tx, Tx{
				Price: tx.Price,
			})
		}

		if okFrom && !okTo {
			intervals[valFrom].tx = append(intervals[valFrom].tx, Tx{
				Price: tx.Price,
			})
		}

		if !okFrom && okTo {
			intervals[valTo].tx = append(intervals[valTo].tx, Tx{
				Price: tx.Price,
			})
		}
	}
	// x arr from
	// y arr to
	return intervals
}

func main() {
	fmt.Println(math.Ceil(1.1))
	//fmt.Println(math.Ceil(1.7))
	a := Fook()
	fmt.Println(a)
}
