package main

func main() {
	//cards := newDeck()
	//cards.saveToFile("mycards")
	cards := newDeckFromFile("mycards1")
	cards.print()
}
