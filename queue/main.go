package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []interface{}
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(item interface{}) {

	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() (interface{}, error) {

	if q.IsEmpty() {
		return nil, errors.New("empty queue")
	}

	val := q.items[0]
	q.items = q.items[1:]
	return val, nil
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return q.items[0], nil
}

//Enqueue
//Dequeue
//Peek

//FIFO
//Task Scheduling
func main() {
	queue := NewQueue()
	queue.Enqueue("first")
	queue.Enqueue("second")
	val, _ := queue.Dequeue()
	fmt.Println(val)
	val, _ = queue.Dequeue()
	fmt.Println(val)

}
