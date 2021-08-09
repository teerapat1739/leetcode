package algorithms

import "fmt"

// Stack represent a LIFO data structure
type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop
func (s *Stack) Pop() int {
	lenStack := len(s.items)
	if lenStack == 0 {
		return 0
	}
	toRemove := s.items[lenStack-1]
	s.items = s.items[:lenStack-1]
	return toRemove
}

func Demo() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
