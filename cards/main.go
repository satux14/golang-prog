package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())
	hand, remainingDeck := deal(cards, 5)

	fmt.Println("Hand:")
	hand.print()

	fmt.Println("Remaining:")
	remainingDeck.print()

	/* Type conversion
	greeting := "Hi there"
	fmt.Println([]byte(greeting))
	*/
}
