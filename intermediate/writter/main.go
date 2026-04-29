package main

import (
	"fmt"
	"io"
	"os"
)

// used to write data to the data stream
// *os.File , os.Stdout , http.ResponseWriter

func main() {
	std := os.Stdout

	writeContent(std, "Writer wrote this \n")

	// file, err := os.Create("sample.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()

	// writeContent(file, "Writer wrote this")

}

func writeContent(w io.Writer, message string) (int, error) {
	n, err := w.Write([]byte(message))

	if err != nil {
		return 0, fmt.Errorf("error while writing: %w", err)
	}

	return n, nil
}
