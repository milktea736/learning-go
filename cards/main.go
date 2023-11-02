package main

var saveFileName string = "my_cards.csv"

func main() {
	cards := newDeck()

	// hand, _ := deal(cards, 5)
	// fmt.Println(hand.toString())

	cards.saveToFile(saveFileName)

	saveCards := readDeckFromFile(saveFileName)
	saveCards.print()
}
