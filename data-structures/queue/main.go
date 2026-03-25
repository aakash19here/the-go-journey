package main

import (
	"fmt"
	"log"
)

type Queue struct {
	Items []int
}

func (q *Queue) Enqueue(item int) {
	q.Items = append(q.Items, item)
}
func (q *Queue) Dequeue() {

	if len(q.Items) < 1 {
		log.Fatal("Cannot dequeue from an empty queue")
	}

	q.Items = q.Items[1:]
}

func (q *Queue) Print() {
	if len(q.Items) < 1 {
		fmt.Println("[]")
	}

	for _, v := range q.Items {
		fmt.Println(v)
	}
}

func main() {
	q := Queue{Items: []int{1, 2, 3, 4}}

	q.Enqueue(5)
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	q.Print()
}
