package main

import (
	"fmt"
	"strings"
)

func main() {

	basket := "Banana, Cat Food, Bread, Avocado"

	Items := RemoveSpacesFromItems(basket)

	totalPrice := totalPriceForItemsInBasket(Items)
	totalPriceInPounds, totalPriceInPence := convertTotalPriceToPoundsAndPence(totalPrice)

	fmt.Println("The total Price of the items in your basket is", totalPriceInPounds, "Pounds and", totalPriceInPence, "Pence")
}

func RemoveSpacesFromItems(basket string) []string {

	Items := strings.Split(basket, ",")

	for item := range Items {
		Items[item] = strings.TrimSpace(Items[item])
	}

	return Items
}

func totalPriceForItemsInBasket(Items []string) int {

	basketPrices := map[string]int{
		"Banana":   50,
		"Cat Food": 90,
		"Bread":    70,
		"Avocado":  150,
	}

	totalPrice := 0

	for _, item := range Items {
		if price, found := basketPrices[item]; found {
			totalPrice += price
		}
	}
	return totalPrice
}

func convertTotalPriceToPoundsAndPence(totalPrice int) (int, int) {

	totalPriceInPounds := totalPrice / 100

	totalPriceInPence := totalPrice - totalPriceInPounds*100

	return totalPriceInPounds, totalPriceInPence

}
