package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "one", "two", "three"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
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
	if handSize > len(d) {
		fmt.Println("Handsize cannot be larger than length of deck!!")
		return nil, nil
	}
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) (deck, error) {
	fmt.Printf("Reading from file : %s\n", fileName)
	content, error := os.ReadFile(fileName)
	if error != nil {
		return nil, error
	}
	return deck(strings.Split(string(content), ", ")), nil
}
