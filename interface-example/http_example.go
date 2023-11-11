package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func request() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	// func ReadAll(r Reader) ([]byte, error) {}

	// Body io.ReadCloser
	// type ReadCloser interface {
	// 	Reader
	// 	Closer
	// }

	// resp.Body implements the io.ReadCloser interface, it must also implement the io.Reader interface.
	// This is why we can pass resp.Body to the io.ReadAll function

	// content, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(content))
	// io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

// implement
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	size := len(bs)
	fmt.Printf("Write  %d of bytes\n", size)
	return size, nil
}
