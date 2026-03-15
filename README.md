## The Golang Journey

- Learn golang using projects
- Will build my own Database , compiler , OS soon

# Latest Commit

**chore: add interface exercise**

Added three Go interface exercises under `intermediate/interfaces/`:

- **exercise 1** (`excercise/main.go`): Implements a custom `Writer` interface with a `ContentWriter` struct that writes a byte slice to stdout using `fmt.Printf`.
- **exercise 2** (`excercise-2/main.go`): Implements an `Incrementer` interface backed by an `IntCounter` type that increments and returns an integer counter over 10 iterations.
- **exercise 3** (`excercise-3/main.go`): Demonstrates interface composition — a `WriterCloser` interface (embedding `Writer` and `Closer`) implemented by `BufferedWriterCloser`, which buffers incoming byte writes and flushes them in 8-byte chunks when `Close()` is called.

# Notes about go

- Everything in go is passed by value by default
