package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Printf("%+v\n", colors)
	printMap(colors)

	fmt.Println("---")

	capitals := make(map[string]string)
	capitals["germany"] = "berlin"
	capitals["vietnam"] = "hanoi"
	capitals["singapore"] = "singapore"

	fmt.Printf("%+v\n", capitals)
	delete(capitals, "singapore")
	fmt.Printf("%+v\n", capitals)

	fmt.Println("---")

	var nilMap map[string]string
	// nilMap["foo"] = "bar" // doesn't work!
	fmt.Printf("%+v\n", nilMap)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Printf("%s => %s\n", k, v)
	}
}
