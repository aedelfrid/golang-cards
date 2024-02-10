package main

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v.", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d) - 1] != "King of Hearts" {
		t.Errorf("Expected last card of King of Hearts, but got %v", d[len(d) - 1])
	}
}

func TestDeal(t *testing.T) {
	cards := newDeck()

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

func TestSaveToFileGetFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()

	deck.saveToFile("_deckTesting")

	deckFromFile := getFromFile("_deckTesting")

	if len(deckFromFile) != 52 {
		t.Errorf("Expected deck size of 52, got %v.", len(deckFromFile))
	}

	os.Remove("_deckTesting")
}