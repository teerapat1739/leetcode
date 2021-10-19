package main

import (
	"fmt"
	"time"
)

func main() {
	txDateTime, err := time.Parse("2006-01-02 15:04:05", "2021-09-21 09:53:23")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(txDateTime)

}
