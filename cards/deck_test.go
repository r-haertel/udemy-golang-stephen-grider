package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected 52 cards, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" || d[len(d)-1] != "King of Clubs" {
		t.Errorf("Seems like the order isn't right. First card: %v. Last card: %v", d[0], d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	d := newDeck()
	d.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards, but got %v", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Spades" || loadedDeck[len(loadedDeck)-1] != "King of Clubs" {
		t.Errorf("Seems like the order isn't right. First card: %v. Last card: %v", loadedDeck[0], loadedDeck[len(d)-1])
	}

	os.Remove("_deckTesting")
}
