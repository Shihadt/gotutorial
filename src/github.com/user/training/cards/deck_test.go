package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Length is not correct %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of spades but got %v", d[0])
	}

	if d[len(d)-1] != "four of Clubs" {
		t.Errorf("error %v", d[len(d)-1])
	}
}

func TestNewDeckFromFileAndSaveToFile(t *testing.T) {
	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")
	loadedDeck := newDeckFromFile("_deckTesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Length expected to be 16 not %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
