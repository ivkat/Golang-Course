package main

func main() {
	cards := newDeck()

	//cards.saveToFile("cartitas.txt")

	//a := newDeckFromFile("cartitas.txt")

	//fmt.Println("cartitas!")
	cards.shuffle()
	cards.print()
}
