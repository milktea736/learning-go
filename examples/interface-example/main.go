// package main

// func main() {

// 	// greet()
// 	// request()
// 	// show()

// 	// fileName := os.Args[1]
// 	// printFileContent(fileName)

// }

package main

import "fmt"

func main() {
	c := make(chan string)

	for i := 0; i < 4; i++ {
		go printString("Hello there!", c)
	}

	for {
		fmt.Println(<-c)
	}
}

func printString(s string, c chan string) {
	fmt.Println(s)
	c <- "Done printing."
}
