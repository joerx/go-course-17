package main

import "fmt"

type bot interface {
	getGreeting(name string) string
}

type englishBot struct {
}

type spanishBot struct {
}

type germanBot struct {
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	gb := germanBot{}
	printGreeting(eb, "Adam")
	printGreeting(sb, "Sue")
	printGreeting(gb, "Klaus")
}

func (eb englishBot) getGreeting(name string) string {
	// Very complicated code to generate an English greeting
	return fmt.Sprintf("Hello %s!", name)
}

func (sb spanishBot) getGreeting(name string) string {
	// Even more complicated spanish greeting code
	return fmt.Sprintf("Hola %s!", name)
}

func (gb germanBot) getGreeting(name string) string {
	// German is the worst
	return fmt.Sprintf("Mahlzeit %s!", name)
}

func printGreeting(b bot, name string) {
	fmt.Println(b.getGreeting(name))
}
