package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter = 0
var wg = sync.WaitGroup{}
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100)
	for range 10 {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}

	wg.Wait()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

func sayHello() {
	fmt.Printf("Hello #%d\n", counter)
	m.RUnlock()
	wg.Done()
}
