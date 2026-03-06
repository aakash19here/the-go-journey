package main

import "fmt"

type person struct {
	name string
	age  int
}

func initPerson() person {
	m := person{
		name: "noname",
		age:  50,
	}
	return m
}

// pointer to person struct
func initPersonWAdd() *person {
	m := person{
		name: "noname",
		age:  50,
	}

	fmt.Printf("initPersonWAdd --> %p\n", &m)
	return &m
}

func main() {
	fmt.Println(initPerson())
	fmt.Printf("%p", initPersonWAdd())
}

/*
	- A function's stack frame goes away after the function returns.
	- So if we return the address of a local value, that value cannot stay only in that stack frame.
	- In this case, Go's escape analysis decides the value escapes and allocates it on the heap.
	- Returning a pointer does not always mean heap allocation, but when a value must outlive the function, heap allocation is a common result.
	- Excess heap allocation can add GC pressure and hurt performance.
*/
