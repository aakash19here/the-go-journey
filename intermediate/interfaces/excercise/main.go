package main

import (
	"fmt"
	"log"
)

type Writer interface {
	Write([]byte) (int, error)
}

type ContentWriter struct{}

func (cw ContentWriter) Write(data []byte) (int, error) {
	n, err := fmt.Printf("%s", string(data))

	return n, err
}

func main() {
	data := []byte{1, 2, 4}
	var w = ContentWriter{}

	i, err := w.Write(data)

	if err != nil {
		log.Fatal("Error printing the bytes ->", err)
	}

	fmt.Println(i)

}
