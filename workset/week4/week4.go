package main

import (
	"fmt"
	"strconv"
)

//PART 4
//Create a function that returns the total for the items in the basket
//Each item now has an item code instead of a price
//look up the price of each item according to its code as returned by the getBarcodeAndPrices function
//NOTE: You may not alter the getBarcodeAndPrices function, EXCEPT to change its name

//PART 4.1:
/*
Given the following basket:
basket := [][]string{
        []string{"Banana", "5001"},
        []string{"Cat Food", "Six Thousand And One"},
        []string{"Bread", "7001"},
        []string{"Avocado", "8001"},
    }
you should return a custom error from your function that does the conversion: https://gobyexample.com/errors
Your application should print this error to the console as "Invalid price for item " and the name of the item
And return the total price of the valid items

PART 4.2:
Given the basket:
basket := [][]string{
        []string{"Banana", "5001"},
        []string{"Cat Food", "6001"},
        []string{"Bread", "7001"},
        []string{"Avocado", "8001"},
        []string{"Yacht", "9001"},
    }
you should return a custom error from the your function that attempts to match the Yacht code
Your application should print this error to the console as "Barcode not found for item X with barcode Y"
where X is the name of the item (Yacht) and Y is the barcode of the item (9001)
*/

//the total should be printed to the console

func main() {
	basket := [][]string{
		[]string{"Banana", "5001"},
		[]string{"Cat Food", "6001"},
		[]string{"Bread", "7001"},
		[]string{"Avocado", "8001"},
	}
	total := getTotal(basket)

	fmt.Println("Total:", total)
}

func convertBarCode(barcodeString string) int {

	barcodeBasket, _ := strconv.Atoi(barcodeString)
	return barcodeBasket
}

func getTotal(basket [][]string) int {

	total := 0
	for _, item := range basket {
		itemName := item[0]
		barcodeString := item[1]
		barcodeBasket := convertBarCode(barcodeString)
		total += getPriceFor(itemName, barcodeBasket)
	}
	return total
}

func getPriceFor(itemName string, barcodeBasket int) int {

	barcodeAndPrices := getBarcodeAndPrices()

	barcodeFound := false
	var price int

	for _, item := range barcodeAndPrices {
		barcode := item[0]
		price = item[1]
		if barcodeBasket == barcode {
			barcodeFound = true
			return price
		}
	}
	if barcodeFound == false {
		fmt.Println("Barcode not found for", itemName)

	}
	return 0

}

// this function is parameterless and retruns a two-dimensional array of ints
func getBarcodeAndPrices() [][]int {
	return [][]int{
		[]int{5001, 50},
		[]int{6001, 90},
		[]int{7001, 70},
		[]int{8001, 150},
	}
}
