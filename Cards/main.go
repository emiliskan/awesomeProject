package main

func main() {

	cards := newDeck()
	hand, remainingDeck := cards.deal(3)
	hand.print()
	remainingDeck.print()

}
