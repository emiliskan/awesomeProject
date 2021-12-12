package main

func main() {

	cards := newDeck()

	err := cards.saveToFile("cards")
	if err != nil {
		return
	}

	cards.shuffle()
	cards.print()
}
