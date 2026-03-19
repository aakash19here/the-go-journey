package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// 50 is a buffer , create channel to create interal datastore to store 50 integers
	// why buffer is needed ? lets say you send more than receiving , there must be a place where you store the additional values
	ch := make(chan int, 50)

	wg.Add(2)

	// receive data from channel
	go func(ch <-chan int) {
		// for i := range ch {
		// 	/*
		// 		- Without a close(ch) ,  for loop introduces a deadlock, because now if the sender is done sending the data , we still keep on reading the values as the loop does not have any exit condition
		// 	*/
		// 	fmt.Println(i)
		// }

		// same like above , but more verbose
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}

		wg.Done()
	}(ch)

	// send data to chan
	go func(ch chan<- int) {
		ch <- 40
		ch <- 42
		ch <- 43
		ch <- 44

		/*
			- Close means that the routine is done sending the values and by that the above receiver can know when to end the for loop
			- When you close the channel make sure you don't send data after that otherwise the compiler will panic
		*/
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
