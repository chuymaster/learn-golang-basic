package main

func main() {
	// cards := deck{newCard(), newCard(), "Ace of Diamonds"}
	// cards = append(cards, "Six of Spades")

	cards := newDeck()
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")

	// cards.print()

	hand, _ := deal(cards, 5)

	hand.print()
	// remainingCards.print()

}
