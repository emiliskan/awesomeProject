package main

import (
	"os"
	"testing"
)

func TestNewDeckLength(t *testing.T) {

	d := newDeck()
	if len(d) != 12 {
		t.Errorf("Excpected 12, but got %v", len(d))
	}

}

func TestNewDeckContains(t *testing.T) {

	d := newDeck()
	if d[0] != "Spades of Ace" {
		t.Errorf("Excpected Spades of Ace in 0 position, but got %v", d[0])
	}

}

func TestSaveDeckFileAndNewDeckFromFIle(t *testing.T) {
	filename := "deck_test"
	d := newDeck()
	err := d.saveToFile(filename)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	deckFromFile := newDeckFromFile(filename)
	if len(deckFromFile) != len(d) {
		t.Errorf("Deck from file dosen't equals to saved deck")
	}

	_ = os.Remove(filename)
}
