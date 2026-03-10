package main

import "fmt"

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34.

func main() {
	n := 10
	a, b := 0, 1

	for range n {
		fmt.Println(a)
		a, b = b, a+b
	}

}
