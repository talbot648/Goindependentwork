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
	questions := []string{"What is 1+1?", "What is 2+2?", "What is 3+3", "What is 7x8", "What is 9X7?", "what number does 10001011 represent?"}
	answers := []int{2, 4, 6, 56, 63, 139}

	total := 0
	qNum := 1

	for i := 1; i <= len(questions); {
		question, indexNumber := getQuestion(questions)
		fmt.Println("Question", qNum, question)

		userGuess := getUserGuess()

		result := checkAnswer(userGuess, indexNumber, answers)
		fmt.Println(result)

		total += changeScore(total, result)
		fmt.Println("Your score is", total)

		questions, answers = removeQuestionAndAnswer(questions, answers, indexNumber)
		qNum++
	}
	fmt.Println("That is the end of the quiz! Your score is", total)

}

func getQuestion(questions []string) (string, int) {

	arrayLength := len(questions)
	indexNumber := rand.Intn(arrayLength)
	question := questions[indexNumber]

	return question, indexNumber
}

func getUserGuess() int {
	var userGuess int
	fmt.Print("Your answer: ")
	fmt.Scanln(&userGuess)
	return userGuess
}

func checkAnswer(userGuess, indexNumber int, answers []int) string {
	if userGuess == answers[indexNumber] {
		return "Correct"

	}
	return "Incorrect"
}

func changeScore(total int, result string) int {
	if result == "Correct" {
		return 100
	} else {
		return -50
	}
}

func removeQuestionAndAnswer(questions []string, answers []int, indexNumber int) ([]string, []int) {
	questions = append(questions[:indexNumber], questions[indexNumber+1:]...)
	answers = append(answers[:indexNumber], answers[indexNumber+1:]...)
	return questions, answers
}
