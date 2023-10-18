package main

import (
	"fmt"
)

func nestedLoops() {
	/*for row := 1; row <= 7; row++ {
		for col := 1; col <= row; col++ {
			fmt.Print("*")
		}
		fmt.Println()

	}*/

	rows := 4
	space := rows - 1

	for j := 1; j <= rows; j++ {
		for i := 1; i <= space; i++ {
			fmt.Print(" ")
		}
		space--
		for i := 1; i <= 2*j-1; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	space = 1
	for j := 1; j <= rows-1; j++ {
		for i := 1; i <= space; i++ {
			fmt.Print(" ")
		}
		space++
		for i := 1; i <= 2*(rows-j)-1; i++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}

func arrays() {

	items := []int{0, 6, 8, 21, 15, 99, -4, 3, -5}
	highestNumber([]int(items))
	numbersGreaterThanTen([]int(items))
	totalOfAllNegativeNumbers([]int(items))
	amountOfNumbersAfterNinetyNine([]int(items))
}

func highestNumber(items []int) {

	biggestNumber := items[0]

	for _, items := range items {

		if biggestNumber < items {

			biggestNumber = items
		}
	}
	fmt.Println("The largest number in the array is", biggestNumber)
}

func numbersGreaterThanTen(items []int) {
	numbersOverTen := 0

	for _, items := range items {

		if items > 10 {

			numbersOverTen++
		}
	}
	fmt.Println("The amount of numbers over ten in the array are", numbersOverTen)
}

func totalOfAllNegativeNumbers(items []int) {
	NegativeNumbersTotal := 0

	for _, item := range items {

		if item < 0 {

			NegativeNumbersTotal = NegativeNumbersTotal + item
		}
	}
	fmt.Println("The total of all negative numbers in the array are", NegativeNumbersTotal)
}

func amountOfNumbersAfterNinetyNine(items []int) {
	NumbersAfter99 := 0
	found99 := false
	for _, items := range items {
		if found99 {
			NumbersAfter99++
		}
		if items == 99 {
			found99 = true
		}

	}
	fmt.Println("The amount of numbers after 99 in the array are", NumbersAfter99)
}
