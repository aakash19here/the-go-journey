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
	- We know that as soon as the function call ends it gets removed from the stack and the active frame points back to main frame
	Where does heap allocation comes into play ?
	- Lets say we return the memory address of a variable created. Once the function is called it would have an address but since the frame is ejected from the stack where does now the address point to ?
	- The answer is heap , whenever go compiler notice that there would be something wrong with the stack popping up , it allocates the heap memory
	- Although using too much heap allocation can cause performance bottlenecks
*/
