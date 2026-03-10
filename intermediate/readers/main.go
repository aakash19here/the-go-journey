package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	fi, err := os.Open("./message.txt")

	if err != nil {
		log.Fatal("Error reading the file")
	}

	str := ""
	for {
		//read 8 bytes of data from the string
		data := make([]byte, 8)

		// n will be always 8 (type int)
		n, err := fi.Read(data)

		if err != nil {
			break
		}

		// eg: Do you h
		data = data[:n]

		// unless and until the index consists \n character keep on adding to the string
		if i := bytes.IndexByte(data, '\n'); i != -1 {

			// from beginning till the last character
			str += string(data[:i])

			//update the data as We don't want character overlap
			data = data[i+1:]
			fmt.Println("read", str)
			str = ""
		}

		//when done append the entire string to the
		str += string(data)

	}

	if len(str) != 0 {
		fmt.Println("read", str)
	}
}
