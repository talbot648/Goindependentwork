package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type car struct {
	Brand    string
	Year     int
	topSpeed int
	colour   string
	price    int
}

func main() {
	cars := initialiseCars()
	userRequest := getUserRequest()

	if userRequest == "top speed" {
		printFastestCar(cars)
	}
	if userRequest == "newest car" {
		printNewestCar(cars)
	}
	if userRequest == "price" {
		printCheapestCar(cars)
	}

}

func initialiseCars() []car {
	car1 := car{"BMW", 2020, 180, "red", 22500}
	car2 := car{"VW", 2008, 135, "Black", 2500}
	car3 := car{"Mercedes", 2015, 165, "White", 12000}
	return []car{car1, car2, car3}
}

func getUserRequest() string {

	options := []string{"top speed", "newest car", "price"}

	for {

		fmt.Println("Your options are", options)
		userRequest := strings.ToLower(askUser("What is the main aspect you are loking for in a car?"))

		for option := range options {

			if userRequest == options[option] {
				return userRequest
			}
		}

		fmt.Println("Invalid request, try again")
		fmt.Println("")
	}
}

func fastestCar(cars []car) car {
	var fastestCar car
	for _, car := range cars {
		if car.topSpeed >= fastestCar.topSpeed {
			fastestCar = car
		}

	}
	return fastestCar
}

func newestCar(cars []car) car {
	var newestCar car
	for _, car := range cars {
		if car.Year >= newestCar.Year {
			newestCar = car
		}

	}
	return newestCar
}

func cheapestCar(cars []car) car {
	cheapestCar := cars[0]
	for _, car := range cars {
		if car.price <= cheapestCar.price {
			cheapestCar = car
		}

	}
	return cheapestCar
}

func printFastestCar(cars []car) {
	fastestCar := fastestCar(cars)
	fmt.Println(fastestCar.Brand, fastestCar.Year, "Is the fastest car with a speed of", fastestCar.topSpeed, "MPH")
}

func printNewestCar(cars []car) {
	newestCar := newestCar(cars)
	fmt.Println(newestCar.Brand, newestCar.Year, "Is the newest car available.")
}

func printCheapestCar(cars []car) {
	cheapestCar := cheapestCar(cars)
	fmt.Println(cheapestCar.Brand, cheapestCar.Year, "Is the cheapest car available costing only", cheapestCar.price, "Pounds")
}

func askUser(prompt string) string {
	fmt.Println(prompt)
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	err := scanner.Err()
	if err != nil {
		panic(err)
	}

	return scanner.Text()
}
