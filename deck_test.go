package main

import (
	"testing"
	"reflect"
	"time"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v.", len(d))
	}
}

func TestDeal(t *testing.T) {
	cards := deck{}

	cardSuits := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	handSize := 8

	hand, remainingCards := deal(cards, handSize)

	if len(hand) != handSize {
		t.Errorf("Expected hand length of 8 but got %v.", len(hand))
	}

	if len(remainingCards) != 52 - handSize {
		t.Errorf("Expected remaining deck of 44 but got %v.", len(remainingCards))
	}
}

func TestShuffle(t *testing.T) {
	handControl := newDeck()
	handRandOne := newDeck()
	handRandTwo := newDeck()

	handRandOne.shuffle()
	time.Sleep(1 *time.Millisecond)
	handRandTwo.shuffle()

	if reflect.DeepEqual(handControl, handRandOne) == true {
		t.Error("Failed to change deck.")
	}

	if reflect.DeepEqual(handRandOne, handRandTwo) == true {
		t.Error("Failed to randomize decks.")
	}

}