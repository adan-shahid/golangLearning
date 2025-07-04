package main

import (
		"fmt"
	   "time")

func main() {
	// var card string 
	//card := newCard()
	// WE ARE DEFINING SLICE TO TYPE STRING
	cards := []string{"Ace of Diamonds", newCard()}
	fmt.Println(cards)
	now := time.Now()
	fmt.Println(now)

}

func newCard() string {
	return "Five of Diamonds"
}	