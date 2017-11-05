package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	dl := len(cardSuites) * len(cardValues)
	d := newDeck()

	if len(d) != dl {
		t.Errorf("Expected deck length %v but got %v", dl, len(d))
	}

	firstCard := fmt.Sprintf("%s of %s",
		cardValues[0],
		cardSuites[0],
	)
	lastCard := fmt.Sprintf("%s of %s",
		cardValues[len(cardValues)-1],
		cardSuites[len(cardSuites)-1],
	)

	if d[0] != firstCard {
		t.Errorf("Expected first card to be %v but was %v", firstCard, d[0])
	}

	if d[len(d)-1] != lastCard {
		t.Errorf("Expected last card to be %v but was %v", lastCard, d[len(d)-1])
	}
}

func TestSaveAndLoadDeck(t *testing.T) {
	fileName := "/tmp/_decktesting"
	cleanup(fileName)

	d := newDeck()

	err := d.writeToFile(fileName)
	if err != nil {
		t.Error(err)
	}

	dl, err := newDeckFromFile(fileName)
	if err != nil {
		t.Error(err)
	}

	if len(d) != len(dl) {
		t.Errorf("Expected deck loaded from file to have len %v but was %v", len(d), len(dl))
	}

	cleanup(fileName)
}

func cleanup(fileName string) {
	err := os.Remove(fileName)
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}
}
