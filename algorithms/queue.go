package algorithms

import "fmt"

// Queue represents a queue
type Queue struct {
	item []int
}

// Enqueue adds integer at the back
func (q *Queue) Enqueue(item int) {
	q.item = append(q.item, item)
}

// Dequeue removes the integer at the front of the queue
// and RETURNNs the removed integer
func (q *Queue) Dequeue() int {
	if len(q.item) == 0 {
		return 0
	}
	toRemove := q.item[0]
	q.item = q.item[1:]
	return toRemove
}

func DemoQueue() {
	myQueue := Queue{}
	myQueue.Enqueue(100)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)
}
