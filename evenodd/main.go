package main

import "fmt"

// comments:
// - no ternary operator
// - is there a better way to generate a range of integers?

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, i := range ints {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}
