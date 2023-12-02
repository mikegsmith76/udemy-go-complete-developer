package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	faces := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}

	for _, suit := range suits {
		for _, face := range faces {
			cards = append(cards, face+" of "+suit)
		}
	}

	return cards
}

func newDeckFromFile(fileName string) deck {
	data, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(data), "\n"))
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)

	for i, _ := range d {
		newPosition := generator.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, d.toByteSlice(), 0666)
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}
