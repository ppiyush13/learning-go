package main

import "testing"

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
