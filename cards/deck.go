package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create a new type of deck
// which is slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Clubs", "Diamond", "Spades", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuite)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}