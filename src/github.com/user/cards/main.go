package main

import (
	"fmt"
)

func main() {
	cards := newDeck()
	//cards.saveToFile("sample.txt")
	//cards := newDeckFromFile("sample.txt")
	//cards.print()
	fmt.Println("Shuffling")
	cards.shuffle()
	cards.print()
}
