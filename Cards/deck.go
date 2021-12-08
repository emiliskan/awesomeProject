package main

import "fmt"

// Deck
type deck []string

func newDeck() deck {

	cardSuits := []string{
		"Spades",
		"Diamonds",
		"Hearts",
	}

	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
	}

	cards := deck{}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}

	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
