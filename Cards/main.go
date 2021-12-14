package main

func main() {

	cards := newDeck()

	err := cards.saveToFile("cards")
	if err != nil {
		return
	}

	hand, _ := deal(cards, 5)
	hand.shuffle()
	hand.print()
}
