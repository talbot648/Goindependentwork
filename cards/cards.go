package main

import "fmt"

func main() {
	cards := ReturnfavouriteCards()
	fmt.Println(cards)
	retrieveCardFromStack()

}

func ReturnfavouriteCards() []int {
	return []int{2, 6, 9}
}

/*func isChosenCardOutOfBounds(slice []int, index int) bool {
	if index < 0 || index >= len(slice) {

		return true
	}

	return false
}*/

func getItem(slice []int, index int) int {

	/*if outOfBounds(slice, index) {

		return -1
	}

	return slice[index]*/

	for i, card := range slice {
		if i == index {
			return card
		}
	}
	return -1

}
func retrieveCardFromStack() {

	cards := []int{1, 2, 4, 1}

	card := getItem(cards, 2)
	fmt.Println(card)
}

/*func setItem() {

}

func exchangeCardInAStack() {
	index := 2

	newCard := 6

	cards := SetItem([]int{1, 2, 4, 1}, index, newCard)

	fmt.Println(cards)
}
*/
