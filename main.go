package main

// import "fmt"

func main() {
	cards := newDeckFromFile("my-cards")

	cards.print()
}