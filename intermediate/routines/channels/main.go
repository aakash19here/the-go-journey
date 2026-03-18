package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // strongly typed channel

	wg.Add(2)

	go func() {
		i := <-ch
		fmt.Println(i)

		wg.Done()
	}()
	go func() {
		i := 42
		ch <- i // pass the copy of i
		i = 40  // receiving go routine won't care if we manipulate the value later
		wg.Done()
	}()

	wg.Wait()

}
