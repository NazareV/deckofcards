package main

func main() {
	//cards := newDeck()

	// hand, remainingcards := deal(cards, 5)

	// hand.print()
	// remainingcards.print()

	//fmt.Println(cards.toString())

	//cards.saveToFile("my_cards")

	cards := newDeckFromFile("my_cards")

	cards.print()
}
