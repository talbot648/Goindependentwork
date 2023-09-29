package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

type Answer struct {
	text      string
	isCorrect bool
}
type QuestionAnswers struct {
	question string
	options  []Answer
}

func main() {
	questions := []QuestionAnswers{
		{
			question: "What is 1+1?",
			options: []Answer{
				{text: "1", isCorrect: false},
				{text: "2", isCorrect: true},
				{text: "3", isCorrect: false},
				{text: "4", isCorrect: false},
			},
		},
		{
			question: "What is 2+2?",
			options: []Answer{
				{text: "2", isCorrect: false},
				{text: "3", isCorrect: false},
				{text: "4", isCorrect: true},
				{text: "5", isCorrect: false},
			},
		},
		{
			question: "What is 7 X 8?",
			options: []Answer{
				{text: "49", isCorrect: false},
				{text: "56", isCorrect: true},
				{text: "64", isCorrect: false},
				{text: "63", isCorrect: false},
			},
		},
	}

	total := 0
	questionNumber := 1

	for i := 1; i <= len(questions); {
		question, indexNumber := getQuestion(questions)

		fmt.Printf("Question %d: %s\n", questionNumber, question.question)
		fmt.Println("Your options are:")
		for _, option := range question.options {
			fmt.Println(option.text)
		}

		userGuess := getUserGuess(question)

		result := checkAnswer(userGuess)
		fmt.Println(result)

		total += addScore(result)
		fmt.Println("Your score is", total)

		questions = removeQuestion(questions, indexNumber)
		questionNumber++
	}
	fmt.Println("That is the end of the quiz! Your score is", total)

}

func getQuestion(questions []QuestionAnswers) (QuestionAnswers, int) {

	arrayLength := len(questions)
	indexNumber := rand.Intn(arrayLength)
	question := questions[indexNumber]

	return question, indexNumber
}

func getUserGuess(question QuestionAnswers) Answer {
	for {
		userGuess := askUser("Your Answer:")
		for _, option := range question.options {
			if userGuess == option.text {
				return option
			}
		}
		fmt.Print("Not an option, retry your answer")
		fmt.Println("")

	}
}

func checkAnswer(userGuess Answer) string {
	if userGuess.isCorrect {
		return "Correct"

	}
	return "Incorrect"
}

func addScore(result string) int {
	if result == "Correct" {
		return 100
	} else {
		return -50
	}
}

func removeQuestion(questions []QuestionAnswers, indexNumber int) []QuestionAnswers {
	return append(questions[:indexNumber], questions[indexNumber+1:]...)
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
