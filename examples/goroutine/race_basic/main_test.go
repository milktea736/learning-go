package main

import "testing"

func TestUpdateMessageWithoutMutex(t *testing.T) {
	msg = "Hello, word!"

	wg.Add(1)
	go updateMessageWithoutMutex("Goodbye")
	go updateMessageWithoutMutex("YoYo")
	wg.Wait()

	if msg != "Goodbye" {
		t.Error("incorrect value")
	}
}

// go test  .  vs go test  -race .
