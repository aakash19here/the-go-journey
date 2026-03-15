package main

import (
	"bytes"
	"fmt"
	"log"
)

type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	//write data
	n, err := bwc.buffer.Write(data)

	if err != nil {
		log.Fatal("Cannot write data")
		return 0, err
	}

	v := make([]byte, 8)

	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)

		if err != nil {
			log.Fatal("Cannot read data")
			return 0, err
		}

		fmt.Println(string(v))
	}

	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)

		_, err := fmt.Println(string(data))

		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("Hello Youtube listeners, this is a test  "))
	wc.Close()

}
