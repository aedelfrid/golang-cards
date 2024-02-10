package main

import "fmt"

func main() {
	startPrompt()

	game := gameSelect()

	dealer := newDeck(game)

	fmt.Println(dealer)
}