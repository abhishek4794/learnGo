package main

import "fmt"

func main() {

	cards := newDeck()
	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()
	deckString := cards.toString()
	fmt.Println(deckString)
	cards.saveToFile("cards.csv")
	//newDeck := newDeckFromFile("cards.csv")
	//newDeck.print()
	//cards.print()
	cards.shuffle()
	cards.print()
}
