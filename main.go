package main

func main() {
	//cards := newDeckFromFile("./deck.dat")
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingCards := deal(cards, 2)

	// hand.print()
	// remainingCards.print()

	// fmt.Println(cards.toString())
	//cards.saveToFile("./deck.dat")
}
