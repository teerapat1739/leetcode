package main

import "fmt"

func main() {
	s := "abc"
	m := make(map[string]bool)
	for _, c := range s {
		m[string(c)] = true
	}
	fmt.Println(m)
	a := m
	delete(a, "a")
	fmt.Println("after", m)
}
