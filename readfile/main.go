package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, f)
}

func usage() {
	fmt.Println("Usage: main.go <filename>")
	os.Exit(1)
}
