package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//spin up a green thread
	// go sayHello()

	var msg = "Hello"

	// tell the wait group that we have another go routine apart from main go routine
	wg.Add(1)

	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)

	msg = "World"
	wg.Wait()
}

// func sayHello() {
// 	fmt.Println("Helllo!!")
// }
