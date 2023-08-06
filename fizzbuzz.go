package main

import (
	"fmt"
)

func fizzBuzz() {
	for count := 1; count <= 100; count++ {
		/*if count%15 == 0 { */
		if count%3 == 0 && count%5 == 0 {
			fmt.Println("fizzbuzz")

		} else if count%3 == 0 {
			fmt.Println("fizz")

		} else if count%5 == 0 {
			fmt.Println("buzz")

		} else {
			fmt.Println(count)
		}
	}

}
