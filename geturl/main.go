package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type stdoutWriter struct {
}

func main() {
	resp, err := http.Get("http://golang.org")
	if err != nil {
		log.Fatal(err)
	}

	// bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("status %s\n", resp.Status)
	fmt.Println("---")

	_, err = io.Copy(stdoutWriter{}, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}

func (st stdoutWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}
