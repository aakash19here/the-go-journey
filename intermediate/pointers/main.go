package main

import "fmt"

/**
-> Pointers as useful in order to share variables from a single place instead of copying it everywhere
-> Also useful when you want to manipulate a variable inside a function
*/

func main() {

	i, j := 42, 1023

	fmt.Println(&i, &j)

	// stores the address of x
	p := &i

	// * returns the value the pointer is pointing to
	// also called dereferencing
	fmt.Println("P is", *p)

	*p = 21

	// pass address only, copying address of i
	squareValue(p)
	fmt.Println("i =", i)

	p = &j
	*p = *p / 32

	fmt.Println("j =", j)

}

func squareValue(v *int) {
	// *v -> value at P
	*v *= *v
}
