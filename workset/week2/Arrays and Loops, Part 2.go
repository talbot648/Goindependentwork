package main

import (
	"fmt"
)

//PART 2
//Instead of a slice, the items have been provided as a single string
//Create a function that
//1. accepts the string as a parameter
//2. calculates the total
//3. returns the total as TWO integers - one for pounds and one for pence

//Note that the cost of each item is:
//Banana 50p
//Cat Food 90p
//Bread 70p
//Avocado £1.50

//https://gobyexample.com/string-functions

//the total should be printed to the console

func main() {
	basket := "Banana, Cat Food, Bread, Avocado"

	fmt.Println("Total:", basket)
}
