package main

import "fmt"

type car struct {
	Brand    string
	Year     int
	topSpeed int
	colour   string
}

func main() {
	car1 := car{"BMW", 2020, 180, "red"}
	car2 := car{"VW", 2008, 135, "Black"}
	car3 := car{"Mercedes", 2015, 165, "White"}
	cars := []car{car1, car2, car3}

	fastestCar := fastestCar(cars)
	fmt.Println(fastestCar.Brand, fastestCar.Year, "Is the fastest car with a speed of", fastestCar.topSpeed)
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
