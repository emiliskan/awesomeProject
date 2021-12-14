package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Deck
type deck []string

func newDeck() deck {

	cardSuits := [3]string{
		"Spades",
		"Diamonds",
		"Hearts",
	}

	cardValues := [4]string{
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

func newDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ERROR:", err)
		return newDeck()
	}

	return strings.Split(string(data), ", ")
}

func (d deck) saveToFile(filename string) error {
	data := []byte(d.toString())
	return ioutil.WriteFile(filename, data, 0666)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	localRand := rand.New(source)

	for i := range d {
		newPos := localRand.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}
