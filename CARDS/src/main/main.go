package main

//var deck int

//card := "Ace of Spades"
func main() {

	cards := newDeck()
	//	cards.print()
	hand, remainingCards := deal(cards, 5)

	hand.print()

	remainingCards.print()
}

func getCard() string {
	return "four of Diamonds"
}
