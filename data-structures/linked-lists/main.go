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
	fmt.Printf("nil, ")
	fmt.Printf("length = %d \n", l.Length)
}

func (l *LinkedList) DecrementLength() {
	l.Length--
}

func (l *LinkedList) DeleteNode(n *Node) {
	if n == nil || l.Head == nil {
		return
	}

	if l.Head == n {
		secondNode := l.Head.Next
		l.Head = secondNode
		l.DecrementLength()
		return
	}

	currentHead := l.Head
	for currentHead != nil {
		if currentHead.Next == n {
			currentHead.Next = n.Next
			l.DecrementLength()
			return
		}

		currentHead = currentHead.Next
	}

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
	list.DeleteNode(Node1)
	list.Display()
}
