package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	var suits [4]string = [4]string{"Spades", "Hearts", "Diamonds", "Clubs"}
	var numbers [13]string
	for i := 0; i < 13; i++ {
		numbers[i] = strconv.Itoa(i + 1)
	}

	cards := deck{}
	for _, s := range suits {
		for _, n := range numbers {
			cards = append(cards, n+" of "+s)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func readDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}
