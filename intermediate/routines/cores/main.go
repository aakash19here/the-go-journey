package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Threads ->", runtime.GOMAXPROCS(-1))
}
