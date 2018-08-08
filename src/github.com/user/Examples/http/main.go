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

		fmt.Println(err.Error())
		os.Exit(1)
	}
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (lw logWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	fmt.Println("Length : ", len(p))
	return len(p), nil
}
