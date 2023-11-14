package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestUpdateMessage(t *testing.T) {
	originalMessage := "A"
	expectedMessage := "B"

	msg = originalMessage
	updateMessage(expectedMessage)

	if msg != expectedMessage {
		t.Errorf("Expected %s, but got %s", expectedMessage, originalMessage)
	}
}

func TestPrintMessage(t *testing.T) {
	message := "A"
	msg = message

	r, w, _ := os.Pipe()
	os.Stdout = w
	var wg sync.WaitGroup

	wg.Add(1)
	printMessage(&wg)
	wg.Wait()

	w.Close()
	bs, _ := io.ReadAll(r)

	output := strings.TrimSpace(string(bs))
	if output != message {
		t.Errorf("Expected %s, but got %s", message, output)
	}
}

func TestMain(t *testing.T) {
	firstMessage := "Hello, universe!"
	secondMessage := "Hello, cosmos!"
	thirdMessage := "Hello, world!"
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()
	_ = w.Close()
	bs, _ := io.ReadAll(r)

	result := string(bs)
	expectedResult := fmt.Sprintf("%s\n%s\n%s\n", firstMessage, secondMessage, thirdMessage)

	if expectedResult != result {
		t.Errorf("Expect %s, but got %s", expectedResult, result)
	}
}
