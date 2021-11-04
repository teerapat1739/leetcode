package main

import (
	"fmt"
	"time"
)

func main() {

	current := time.Now()
	f,_ := time.LoadLocation("UTC")
	a := current.In(f)

	fmt.Println(a.String())
}
