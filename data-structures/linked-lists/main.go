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

func (l *LinkedList) InsertAtEnd(n *Node) {
	// copy the old head node
	secondNode := l.Head

	//make the new node head
	l.Head = n

	// attach the second node to the head of new node as next node
	l.Head.Next = secondNode

	//update the length of linked list
	l.Length++
}

func (l *LinkedList) Display() {
	current := l.Head    // Start from the head
	for current != nil { // Loop until the current node is nil (end of list)
		fmt.Printf("%d -> ", current.Data) // Print the current node's data
		current = current.Next             // Move to the next node
	}
	fmt.Println("nil") // Indicate the end of the list
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

	list.InsertAtEnd(Node1)
	list.InsertAtEnd(Node2)
	list.InsertAtEnd(Node3)

	list.Display()
}
