package main

import "fmt"

func main() {
	cards := newDeck()

	// hand, remainingcards := deal(cards, 5)

	// hand.print()
	// remainingcards.print()

	fmt.Println(cards.toString())
}
