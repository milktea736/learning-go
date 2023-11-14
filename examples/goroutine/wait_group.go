package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func run_wait_group() {
	words := []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
	}
	var wg sync.WaitGroup
	wg.Add(len(words))

	for i, v := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, v), &wg)
	}

	wg.Wait()

	wg.Add(1)
	printSomething("Done", &wg)
}
