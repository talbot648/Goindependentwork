package main

import (
	"fmt"
	"strings"
)

//PART 5
//Adding items to a basket, getting the code, getting the price, and totalling the basket
//You have been provided a string (items) that you may not alter directly (that is, you can't change the contents on line 34 just to make your life easier...)

//Step 1 - create a function that gets the item code
//The function should accept one parameter - a string (e.g 'Banana' or 'Cat Food') // for in range of something, call the function
//and returns a slice of type *string* that contains the item and the item code
//Item names: Banana, Cat Food, Bread, Avocado
//Item codes: 999999, 33303333, 88888, 4444444

//Step 2
//**Append** the returned slice to the basket
//https://yourbasic.org/golang/append-explained/

//Step 3
//Once all items have been added to the basket, total the item prices of the basket
//This should be in a function that accepts the entire basket as a parameter and returns a single int
//Each Item has a price
//Item price: 99p, 1.76, 2.99, 67p

//HINT: You should write another function that returns the price when given a code, rather than putting all the logic in the basket totalling function...

//the total should be printed to the console

func main() {

	//var items string = "banana"
	//var items string = "baNana"
	//var items string = "baNana,brEad"
	//var items string = "    baNana ,breAD "
	var items string = " banana, Banana, BanANa,   Cat FOOD, Cat Food  ,    Bread, AvoCADo   ,,,"

	var basket = [][]string{}

	itemsAndCode := returnItemAndItemCode(items)
	basket = append(basket, itemsAndCode...)
	fmt.Println(basket)
	total := getTotalPrice(basket)
	fmt.Println("Total:", total, "Pence")
}

func returnItemAndItemCode(items string) [][]string {
	Items := extractItems(items)
	var itemsAndCode [][]string
	for _, item := range Items {
		itemCode := getItemCode(item)
		itemsAndCode = append(itemsAndCode, []string{item, itemCode})
	}
	return itemsAndCode

}

func extractItems(items string) []string {

	Items := strings.Split(items, ",")
	var cleanedItems []string

	for index := range Items {
		cleanedItem := strings.TrimSpace(Items[index])
		if cleanedItem != "" {
			cleanedItem = strings.ToLower(cleanedItem)
			cleanedItem = strings.Title(cleanedItem)
			cleanedItems = append(cleanedItems, cleanedItem)
		}
	}
	return cleanedItems
}

func getItemCode(item string) string {
	var itemCode string
	if item == "Banana" {
		itemCode = "999999"
		return itemCode
	}
	if item == "Cat Food" {
		itemCode = "33303333"
		return itemCode
	}
	if item == "Bread" {
		itemCode = "88888"
		return itemCode
	}
	if item == "Avocado" {
		itemCode = "4444444"
		return itemCode
	}
	fmt.Println("There are no matching codes for item", item)
	return itemCode

}

func getTotalPrice(basket [][]string) int {
	total := 0
	for item := range basket {
		barcode := basket[item][1]
		total += getPriceFromCode(barcode)
	}
	return total

}

func getPriceFromCode(barcode string) int {
	var total int
	if barcode == "999999" {
		total += 99
		return total
	}
	if barcode == "33303333" {
		total += 176
		return total
	}
	if barcode == "88888" {
		total += 299
		return total

	}
	if barcode == "4444444" {
		total += 67
		return total

	}
	return total

}
