package main

import (
	"fmt"
)

// data -> hash fn -> get index of element

/*
- What we need ?
- Hashtable structure
- Bucket structure

Methods
- Insert for hashtable and bucket
- Search for hashtable and bucket
- Delete for hashtable and bucket

Hash func
Init
*/

// size of hashtable array
const SIZE int = 7

type Hashtable struct {
	array [SIZE]*Bucket
}

type Bucket struct {
	Head   *BucketNode
	Length int
}

type BucketNode struct {
	Next *BucketNode
	Data string
}

func main() {
	hashTable := InitHashTable()

	list := []string{"ERIC", "RANDY", "KENNY", "STAN", "TOKEN"}

	for _, v := range list {
		hashTable.Insert(v)
	}

	hashTable.Delete("STAN")
	fmt.Println("STAN", hashTable.Search("STAN"))
	fmt.Println("KENNY", hashTable.Search("KENNY"))

}

func InitHashTable() *Hashtable {

	result := &Hashtable{}

	for i := range result.array {
		// attach bucket to each of the array element
		result.array[i] = &Bucket{}
	}

	return result
}

// hash table methods
func (h *Hashtable) Insert(data string) {
	index := hash(data)

	h.array[index].Insert(data)
}

func (h *Hashtable) Delete(data string) {
	index := hash(data)

	h.array[index].Delete(data)
}

func (h *Hashtable) Search(data string) bool {
	index := hash(data)

	return h.array[index].Search(data)
}

// bucket methods
func (b *Bucket) Insert(data string) {
	if !b.Search(data) {
		newNode := &BucketNode{Data: data}
		newNode.Next = b.Head
		b.Head = newNode
	} else {
		fmt.Println(data, "already exists")
	}
}

func (b *Bucket) Search(data string) bool {
	currentNode := b.Head

	for currentNode != nil {
		if currentNode.Data == data {
			return true
		} else {
			currentNode = currentNode.Next
		}
	}

	return false
}

func (b *Bucket) Delete(data string) {
	if b.Head.Data == data {
		b.Head = b.Head.Next
		return
	}

	prevNode := b.Head
	for prevNode.Next != nil {
		if prevNode.Next.Data == data {
			prevNode.Next = prevNode.Next.Next
		}
		prevNode = prevNode.Next
	}
}

func hash(data string) int {
	sum := 0

	for _, v := range data {
		sum += int(v)
	}
	return sum % SIZE
}
