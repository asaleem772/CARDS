package main

//import (
//	"fmt"
//)

func main() {

	cards := newDeck()
	cards.shuffle()
	cards.print()
	//	cards.saveToFile("cards")
	//
	//	newcards := newDeckFromFile("scards")
	//	newcards.print()

}

func getCard() string {
	return "four of Diamonds"
}
