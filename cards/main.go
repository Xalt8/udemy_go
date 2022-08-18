package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades" //Only use walrus operator when creating a new variable
	// card = "Five of Diamonds" // can re-assign variable with = operator

	cards := newDeckFromFile("my_cards")
	cards.print()

	// hand, reamingCards := deal(cards, 5)

}
