package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "there", "how", "are", "you?"}
	fmt.Println(mySlice)
	updateSlice(mySlice, "Hello")
	fmt.Println(mySlice)
}

func updateSlice(s []string, by string) {
	s[0] = by
}
