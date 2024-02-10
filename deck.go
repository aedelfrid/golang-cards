package main

import (
	"fmt"
	//"os"
	//"strings"
	"math/rand"
	"time"
)

type deck []card

type card struct {
	name string
	rank int
}

func newDeck(p playset) deck {

	cards := deck{}

	cardSuits := []string{
		"Spades",
		"Clubs",
		"Diamonds",
		"Hearts",
	}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
	}

	for _, suit := range cardSuits {
		for i, value := range cardValues {
			cards = append(
				cards,
				card{value + " of " + suit, p.cardValues[i]},
			)

		}
	}

	return cards
}

func (d deck) deal(handSize int) deck {
	d = d[:handSize]
	return d[handSize:]
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixMicro())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// func (d deck) toString() string {
// 	return strings.Join([]string(d), ",")
// }

// func (d deck) saveToFile(filename string) error {
// 	return os.WriteFile(filename, []byte(d.toString()), 0666)
// }

// func getFromFile(filename string) deck {
// 	bs, err := os.ReadFile(filename)

// 	if err != nil {
// 		fmt.Println("Error:",err)

// 		return newDeck()
// 	}

// 	return deck(strings.Split(string(bs), ","))

// }
