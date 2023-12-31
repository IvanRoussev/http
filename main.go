package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}



func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}


func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes: ", len(bs))

	return len(bs), nil
}

func anotherRequest() {
	resp, err := http.Get("http://amazon.com")

	if err != nil {
		fmt.Println(err)
	}

	lw := logWriter{}
	
	io.Copy(lw, resp.Body)
}