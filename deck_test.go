package main

import (
	"os"
	"testing"
)

/*
	1. Length of new deck
	2. Assert first card
	3. Assert last card
*/

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 20 {
		t.Errorf("Expected deck of length of 20, but got %v", len(d))
	}
}

func TestFirstCard(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
}

func TestLastCard(t *testing.T) {
	d := newDeck()

	if d[len(d)-1] != "Five of Hearts" {
		t.Errorf("Expected last card of Five of Hearts, but got %v", d[len(d)-1])
	}
}

func TestSaveTofile(t *testing.T) {
	os.Remove("./tmp/_deck.dat")

	deck := newDeck()
	deck.saveToFile("./tmp/_deck.dat")

	deck2 := newDeckFromFile("./tmp/_deck.dat")

	if len(deck2) != 20 {
		t.Errorf("Expected deck of length of 20, but got %v", len(deck2))
	}
}
