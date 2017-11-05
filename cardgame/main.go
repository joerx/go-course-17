package main

func main() {
	deck := newDeck().shuffle()
	hand, deck := deal(deck, 5)
	hand.print()
	hand.writeToFile("./deck.txt")
}
