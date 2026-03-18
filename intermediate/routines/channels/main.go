package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int) // strongly typed channel
	for range 5 {
		wg.Add(2)
		go func() {
			i := <-ch
			fmt.Println(i)

			wg.Done()
		}()
		go func() {
			i := 42
			ch <- i // pass the copy of i
			wg.Done()
		}()
	}
	wg.Wait()

}
