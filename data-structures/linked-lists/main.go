package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func (l *LinkedList) InsertAtStart(n *Node) {
	// copy the old head node
	secondNode := l.Head

	//make the new node head
	l.Head = n

	// attach the second node to the head of new node as next node
	l.Head.Next = secondNode

	//update the length of linked list
	l.Length++
}

func (l LinkedList) Display() {
	current := l.Head    // Start from the head
	for current != nil { // Loop until the current node is nil (end of list)
		fmt.Printf("%d -> ", current.Data) // Print the current node's data
		current = current.Next             // Move to the next node
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList{}

	Node1 := &Node{
		Data: 1,
	}
	Node2 := &Node{
		Data: 2,
	}
	Node3 := &Node{
		Data: 3,
	}

	list.InsertAtStart(Node1)
	list.InsertAtStart(Node2)
	list.InsertAtStart(Node3)

	list.Display()
}
