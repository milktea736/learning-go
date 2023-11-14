package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessageWithoutMutex(s string) {
	defer wg.Done()

	msg = s
}

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "Hello, world!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("A", &mutex)
	go updateMessage("B", &mutex)
	wg.Wait()

	fmt.Println(msg)
}

// go run -race main.go
