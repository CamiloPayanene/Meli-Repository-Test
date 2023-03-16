package main

import "fmt"



func main(){
	cards := []string{newCard(), "ace of diamonds"}
	cards = append(cards, "six of trebols")

	for i, card := range cards{
		fmt.Println(i,card)
	}
}

func newCard() string {

	return "Five of pikas"
}