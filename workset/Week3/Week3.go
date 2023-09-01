package main

import (
	"fmt"
	"strconv"
	"strings"
)

//PART 3
//Create a new, well-named function that returns the total as a single integer for the items in the basket
//Each item now has its own price, expressed as a *string*. You cannot add strings together, so you will need to find out how to convert it.
//Note: the basket is now a two-dimensional array - we have an array that contains arrays
//https://www.tutorialspoint.com/go/go_multi_dimensional_arrays.htm

/* EXTRA CREDIT:
1. Solve for this basket:
basket := [][]string{
		[]string{" Banana ", " 50  "},
		[]string{"Cat Food", "    90"},
		[]string{"Bread    ", "70   "},
		[]string{"   Avocado ", "  150"},
	}
2. Implement error handling:
basket := [][]string{
		[]string{" Banana ", "50"},
		[]string{"Cat Food", "    90"},
		[]string{"Bread    ", "70   "},
		[]string{"   Avocado ", "  one pound fifty"},
        []string{"Yacht", "150,000.99"},
	}
*/

//the total should be printed to the console

func main() {

	basket := [][]string{
		[]string{"Banana", "50"},
		[]string{"Cat Food", "90"},
		[]string{"Bread", "70"},
		[]string{"Avocado", "150"},
	}
	basketWithSpaces := [][]string{
		[]string{" Banana ", " 50  "},
		[]string{"Cat Food", "    90"},
		[]string{"Bread    ", "70   "},
		[]string{"   Avocado ", "  150"},
	}

	/*invalidbasket := [][]string{
		[]string{" Banana ", "50"},
		[]string{"Cat Food", "    90"},
		[]string{"Bread    ", "70   "},
		[]string{"   Avocado ", "  one pound fifty"},
		[]string{"Yacht", "150,000.99"},
	}*/
	RemoveSpacesFromBasket(basketWithSpaces)
	fmt.Println(basket)
	fmt.Println("The total Price of the items in your basket is", getTotal(basket))
	fmt.Println("The total Price of the items in your basket is", getTotal(basketWithSpaces))
	//fmt.Println("The total Price of the items in your basket is", getTotal(invalidbasket))

}

func RemoveSpacesFromBasket(basket [][]string) [][]string {

	for i := range basket {
		basket[i][1] = strings.TrimSpace(basket[i][1])
	}

	return basket
}

func convertItemPrice(priceString string) int {

	price, _ := strconv.Atoi(priceString)
	return price

}

func getTotal(basket [][]string) int {
	total := 0

	for _, item := range basket {
		priceString := item[1]
		price := convertItemPrice(priceString)
		total += price
	}
	return total
}
