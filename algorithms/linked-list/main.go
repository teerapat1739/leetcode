package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) insert(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) search(value int) bool {

	currentNode := l.head
	for currentNode != nil {
		if currentNode.data == value {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (l *linkedList) delete(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l linkedList) print() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 11}
	node5 := &node{data: 7}
	node6 := &node{data: 2}
	mylist.insert(node1)
	mylist.insert(node2)
	mylist.insert(node3)
	mylist.insert(node4)
	mylist.insert(node5)
	mylist.insert(node6)
	mylist.print()

	mylist.delete(100)
	mylist.delete(2)
	mylist.delete(18)
	mylist.print()

}
