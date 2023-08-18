package main

import (
	"fmt"
)

//PART 1
//Create a function that returns the total (an int) for the all items in the basket
//Note that the cost of each item is:
//Banana 50p
//Cat Food 90p
//Bread 70p
//Avocado £1.50

//the total should be printed to the console

func main() {
	basket := []string{"Banana", "Cat Food", "Bread", "Avocado"}

	totalPrice := totalPriceForItemsInBasket(basket)
	fmt.Println("The total price for the items in the basket is", totalPrice, "P")

}

func totalPriceForItemsInBasket(basket []string) int {

	basketPrices := map[string]int{
		"Banana":   50,
		"Cat Food": 90,
		"Bread":    70,
		"Avocado":  150,
	}

	totalPrice := 0

	for _, item := range basket {
		if price, found := basketPrices[item]; found {
			totalPrice += price
		}
	}
	return totalPrice

}
