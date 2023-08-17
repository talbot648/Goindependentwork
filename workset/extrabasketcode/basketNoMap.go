package main

import (
	"fmt"
)

func main() {
	basket := []string{"Banana", "Cat Food", "Bread", "Avocado", "Bread"}

	totalPrice := totalPriceForItemsInBasket(basket)

	fmt.Println("The total price is", totalPrice, "Pence")
}

func totalPriceForItemsInBasket(basket []string) int {

	totalPrice := 0

	for _, item := range basket {
		if item == "Banana" {
			totalPrice += 50
		} else if item == "Cat Food" {
			totalPrice += 90
		} else if item == "Bread" {
			totalPrice += 70
		} else if item == "Avocado" {
			totalPrice += 150
		}
	}

	return totalPrice
}
