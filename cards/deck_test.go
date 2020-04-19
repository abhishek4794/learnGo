package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 12 {
		t.Error("Expected deck length of 12 but got ", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Error("Expected first card to be Ace of Clubs but got ", d[0])
	}

}

func TestSaveToDeckNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 12 {
		t.Error("Expected deck from file length of 12 but got ", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
