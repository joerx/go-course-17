package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// A type representing a deck of playing cards
type deck []string

var cardSuites = []string{
	"\u2660",
	"\u2666",
	"\u2665",
	"\u2663",
}
var cardValues = []string{
	"Two",
	"Three",
	"Four",
	"Five",
	"Six",
	"Seven",
	"Eight",
	"Nine",
	"Ten",
	"Jack",
	"Queen",
	"King",
	"Ace",
}

// Creates a new full deck of playing cards in order of card values and suits.
func newDeck() deck {
	cards := deck{}
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}
	return cards
}

// Read deck from file at given path. Returns an error when reading the source file failed for
// any reason.
func newDeckFromFile(fileName string) (deck, error) {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	s := strings.Split(string(bs), "\n")
	return deck(s), nil
}

// Print the card deck
func (d deck) print() {
	fmt.Println(d.toString())
}

// Writes deck into a file, fileName is either absolute or relative to working dir
func (d deck) writeToFile(fileName string) error {
	b := []byte(d.toString())
	return ioutil.WriteFile(fileName, b, 0644)
}

// Returns a string representation of a deck by joining all cards in the deck with a \n char
func (d deck) toString() string {
	return strings.Join(d, "\n")
}

// Shuffle the deck, returns a shuffled copy of the deck. Internally using pseudo-random numbers
// initialized with the current system time
func (d deck) shuffle() deck {
	rnd := mkRand()
	sd := d.copy()

	for i := range d {
		n := rnd.Intn(len(sd))
		sd[i], sd[n] = sd[n], sd[i]
	}
	return sd
}

// Deals number given by handSize of cards given from deck and returns the hand and the remaining
// deck of cards.
func deal(d deck, handsize int) (deck, deck) {
	hand := d[:handsize]
	deck := d[handsize:]
	return hand, deck
}

func mkRand() *rand.Rand {
	source := rand.NewSource(time.Now().UnixNano())
	return rand.New(source)
}

func (d deck) copy() deck {
	sd := make(deck, len(d))
	copy(sd, d)
	return sd
}
