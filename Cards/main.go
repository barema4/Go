package main
import "fmt"


func main()  {
	// var cards string = "Ace of shades"
	cards := newDeck()
	
	fmt.Println(cards.toString())
}
