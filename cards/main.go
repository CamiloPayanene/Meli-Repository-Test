package main

func main() {

	cards := deck{newCard(), "ace of diamonds"}
	cards = append(cards, "six of trebols")

	cards.print()
}

func newCard() string {

	return "Five of pikas"
}
