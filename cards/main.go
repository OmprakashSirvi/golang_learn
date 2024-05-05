package main

import (
	"fmt"
	"os"
)

func main() {
	cards, error := newDeckFromFile("onHand1")
	if error != nil {
		fmt.Println("Error : ", error)
		os.Exit(1)
	}
	cards.print()
}
