package main

import (
	"fmt"
	"log"
)

type Stack struct {
	Items []int
}

func (s *Stack) Push(data int) {
	s.Items = append(s.Items, data)
}

func (s *Stack) Pop() {
	if len(s.Items) < 1 {
		log.Fatal("Cannot Pop from an empty stack")
		return
	}
	s.Items = s.Items[:len(s.Items)-1]

}

func (s *Stack) Print() {

	if len(s.Items) < 1 {
		fmt.Println("[]")
	}

	for _, v := range s.Items {
		fmt.Println(v)
	}
}

func main() {
	stack := Stack{
		Items: []int{1, 2, 3},
	}

	stack.Push(10)
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()

	stack.Print()

}
