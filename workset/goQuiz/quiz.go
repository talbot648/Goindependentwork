package main

import (
	"fmt"
	"math/rand"
)

// Plan: ask user a question
// user inputs an answer
// if question is correct then add 100 score
// if wrong deduct 50 score
// store score

func main() {
	questions := []string{"What is 1+1?", "What is 2+2?", "What is 3+3"}

	total := 0
	var userGuess int

	for i := 1; i <= len(questions); i++ {
		question, indexNumber := getQuestion(questions)

		fmt.Println(question)
		fmt.Scanln(&userGuess)

		result := checkAnswer(userGuess, indexNumber)
		fmt.Println(result)

		if result == "Correct" {
			total += addScore(total)
		}
		fmt.Println("Your score is", total)

	}
}

func getQuestion(questions []string) (string, int) {

	arrayLength := len(questions)
	indexNumber := rand.Intn(arrayLength)
	question := questions[indexNumber]

	return question, indexNumber
}

func checkAnswer(userGuess, indexNumber int) string {
	answers := []int{2, 4, 6}
	if userGuess == answers[indexNumber] {
		return "Correct"

	}
	return "Incorrect"
}

func addScore(total int) int {
	return 10

}
