package main

import (
	"io"
	"os"
)

func printFileContent(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, file)
}
