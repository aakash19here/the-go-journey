package main

import "fmt"

func main() {
	n := IntCounter(0)
	var inc Incrementer = &n

	for range 10 {
		fmt.Println(inc.Increment())
	}

}

type IntCounter int

type Incrementer interface {
	Increment() int
}

func (i *IntCounter) Increment() int {
	*i += 1
	return int(*i)
}
