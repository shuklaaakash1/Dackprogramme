package main

func main() {

	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.toString()
	// cards.saveToFile("my_card")
	// // var card string = "Ace of Spade"
	// cards := newDeck()
	// hand, remainingCardd := deal(cards, 5)
	// hand.print()
	// remainingCardd.print()
	// // greeting := "hi there!"
	// // fmt.Println([]byte(greeting))
	//
	cards := newDeck()
	cards.shuffle()
	cards.print()

}
