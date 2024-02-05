package main

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 8)

	hand.print()
	remainingDeck.print()
}