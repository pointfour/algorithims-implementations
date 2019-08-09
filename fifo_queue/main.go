package main

import "fmt"

func main() {
	var queue StringQueue
	queue.Enqueue("Hello World")
	fmt.Println(queue.Dequeue())
	queue.Enqueue("test1")
	queue.Enqueue("test2")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	queue.Enqueue("test1")
	queue.Enqueue("test2")
	fmt.Println(queue.Dequeue())
	queue.Enqueue("test3")
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
}

type node struct {
	Value string
	Next  *node
}

//StringQueue is a queue of strings
type StringQueue struct {
	first *node
	last  *node
	Size  int
}

//Enqueue enqueues a string
func (q *StringQueue) Enqueue(Value string) {
	oldLast := q.last
	newLast := node{}
	q.last = &newLast
	q.last.Value = Value
	q.last.Next = nil
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.Next = q.last
	}
	q.Size++
}

//Dequeue returns the most recent Value from the queue
func (q *StringQueue) Dequeue() string {
	Value := q.first.Value
	q.first = q.first.Next
	q.Size--
	if q.IsEmpty() {
		q.last = nil
	}
	return Value
}

//IsEmpty returns true if the queue is empty
func (q *StringQueue) IsEmpty() bool {
	return q.first == nil
}
