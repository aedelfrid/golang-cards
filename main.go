package main

import "fmt"

func main() {
	startPrompt()

	game := gameSelect()

	dealer := newDeck(game.cardValues)

	player1 := dealer.deal(game.handSize)

	fmt.Println(player1)
}