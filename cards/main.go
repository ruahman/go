package main

import "fmt"

func main() {
	// var card string = "ace of spades"
	card := "ace of spades"
	fmt.Println(card)
	fmt.Println(newCard())
}

func newCard() string {
	return "Five of dimonds"
}
