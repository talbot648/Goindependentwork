package main

import (
	"fmt"
	"math/rand"
)

func crapsGame() {

	for i := 1; i <= 10; i++ {
		diceOne := rollSingle()
		diceTwo := rollSingle()
		fmt.Println("Your roll is", diceOne, diceTwo)
		rollName(diceOne, diceTwo)
	}
}

func rollSingle() int {
	return rand.Intn(6) + 1
}

func rollName(diceOne int, diceTwo int) {
	total := diceOne + diceTwo
	if diceOne == diceTwo {
		printRollNameForDouble(total)
	} else {

		if total == 3 {

			fmt.Println("Ace Deuce")
		}

		if total == 4 {

			fmt.Println("Easy Four")
		}
		if total == 5 {

			fmt.Println("Five(Fever Five)")
		}
		if total == 6 {

			fmt.Println("Easy Six")
		}
		if total == 7 {

			fmt.Println("Natural/Seven Out")
		}
		if total == 8 {

			fmt.Println("Easy Eight")
		}
		if total == 9 {

			fmt.Println("Nine(Nina)")
		}
		if total == 10 {

			fmt.Println("Easy Ten")
		}
		if total == 11 {

			fmt.Println("Yo(Yo-leven)")
		}
		if total == 12 {

			fmt.Println("Boxcars/Midnight")
		}
	}
}

func printRollNameForDouble(total int) {

	if total == 2 {

		fmt.Println("Snake Eyes")
	}

	if total == 4 {

		fmt.Println("Hard Four")
	}
	if total == 6 {

		fmt.Println("Hard Six")
	}
	if total == 8 {

		fmt.Println("Hard Eight")
	}
	if total == 10 {

		fmt.Println("Hard Ten")
	}
}
