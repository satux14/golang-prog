package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	remain := d[handSize:]

	return hand, remain
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
