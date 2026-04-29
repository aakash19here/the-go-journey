package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	// Instead of reading through file we can also use strings.NewReader which implements io.Reader
	// file, err := os.Open("sample.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	r := strings.NewReader("Hello World !!!")

	count, err := readFileContent(r)

	if err != nil {
		panic(err)
	}

	fmt.Println("Count is", count)
}

func readFileContent(reader io.Reader) (int, error) {
	buffer := make([]byte, 8)
	count := 0

	for {
		n, err := reader.Read(buffer)

		if n > 0 {
			count += n
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
	}

	return count, nil
}
