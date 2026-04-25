package main

import (
	"fmt"
	"sync"
)

type SessionTracker struct {
	mu             sync.Mutex
	activeSessions int
}

func (st *SessionTracker) AddSession() {
	st.mu.Lock()         // lock the mutex before accessing the state
	defer st.mu.Unlock() // unlock when its done
	st.activeSessions++
}

func main() {
	tracker := &SessionTracker{}

	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			tracker.AddSession()
		}()
	}

	wg.Wait()

	fmt.Printf("Expected 1000 active sessions, got: %d\n", tracker.activeSessions)
}
