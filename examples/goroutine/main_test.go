package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSom(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add((1))

	go printSomething("hello", &wg)

	wg.Wait()

	_ = w.Close()

	results, _ := io.ReadAll(r)
	output := string(results)

	os.Stdout = stdOut

	if !strings.Contains(output, "hello") {
		t.Errorf("Expect to find hello, but not found")
	}
}
