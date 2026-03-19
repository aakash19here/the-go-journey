package main

import (
	"fmt"
	"time"
)

const (
	info  = "INFO"
	warn  = "WARNING"
	fatal = "FATAL"
	debug = "DEBUG"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logChal = make(chan logEntry, 50)

/*
- Strongly type to a struct with no fields
- Struct with no fields is unique as it requires no mem allocations
- it can't send send any data through except for the fact that the message was sent or received
- this is called as signal only channel
- letting the receiving side know that there was a message sent over the channel
- you can also use bool instead of a empty struct but that requires a mem allocation
*/
var doneCh = make(chan struct{})

func main() {

	// close channel when it's done
	// defer func() {
	// 	close(logChal)
	// }()

	/*
		- without closing (close(ch)) or using select we are not gracefully shutting down the program
		- the main program gets a go routine and then works on the remaining code
		- if at any time the code terminates , we need a way to terminate the routine as well or we can get memory leaks
	*/

	go logger()

	logChal <- logEntry{time: time.Now(), severity: info, message: "App started running on 3000"}
	logChal <- logEntry{time: time.Now(), severity: warn, message: "OPENAI API KEY MISSING"}

	/*
		- hacky way to give go routine time to finish
		- if not given program terminates as soon as main function is done with all of the work
		- concrete way to do this is using Waitgroup
	*/
	time.Sleep(100 * time.Millisecond)

	// passing empty struct lets the compiler know that we are done sending the data
	doneCh <- struct{}{}

}

func logger() {
	// for entry := range logChal {
	// 	fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02 15:04:05"), entry.severity, entry.message)
	// }

	/*
		- Alternative to above block
	*/
	for {
		select {
		case entry := <-logChal:
			fmt.Printf("%v - [%v] %v\n", entry.time.Format("2006-01-02 15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}
	}
}
