package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Hangman tasks
//Choose a random word(from a given list or a csv file)
//split each letter into slices
//create a user guess input
//

func main() {

	word := getWord()
	userGuess := askUser("Guess")
	guessCount := 0

	letterArray := strings.Split(word, ".")
	for _, letter := range letterArray {
		if userGuess == letter {
			fmt.Println("correct")
		} else {
			fmt.Println("incorrct")
			guessCount++
		}
	}

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

func getWord() string {
	words := []string{"c.h.i.l.d", "d.o.g", "a.n.i.m.a.l"}
	return words[0]
}
