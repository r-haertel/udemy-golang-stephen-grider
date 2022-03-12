package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingSize := deal(cards, 3)
	fmt.Println(hand, remainingSize)
}
