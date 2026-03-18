package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	// receive only routine
	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		fmt.Println(<-ch)

		wg.Done()
	}(ch)

	//send only routine
	go func(ch chan<- int) {
		i := 42
		ch <- i
		ch <- 27
		wg.Done()
	}(ch)

	wg.Wait()
}
