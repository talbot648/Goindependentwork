package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

// have a userinput Rock paper or scissors
// have a cpu choose random of the 3
// have a scoreboard to track the scores
// potentially have the option to play to a certain score given by the user or infinitely (or when users says quit)

func main() {
	var userScore int
	var opponentScore int
	fmt.Println("rock paper scissors Game!")
	total := askUser("How many rounds do you want to play?")
	totalInt, err := strconv.Atoi(total)
	if err != nil {
		fmt.Println("Error input")
	}
	fmt.Println("Enter rock, paper or scissors")

	for i := 0; i < totalInt; i++ {
		userChoice := askUser("What is your Choice? ")
		opponentChoice := getOppnentsChoice()
		result := getRoundWinner(userChoice, opponentChoice)
		userScore, opponentScore = addScore(userScore, opponentScore, result)

		fmt.Println("")
		fmt.Println("The scores are", userScore, "-", opponentScore)

	}
	fmt.Println("That is the end of the game!")
	if userScore > opponentScore {
		fmt.Println("You win!")
	} else if userScore < opponentScore {
		fmt.Println("Opponent wins")
	} else {
		fmt.Println("Its a draw")
	}
}

func getOppnentsChoice() string {
	choices := []string{"rock", "paper", "scissors"}
	indexNumber := rand.Intn(len(choices))
	opponentChoice := choices[indexNumber]

	return opponentChoice
}

func getRoundWinner(userChoice string, opponentChoice string) string {
	var result string
	if userChoice == "rock" && opponentChoice == "paper" || userChoice == "paper" && opponentChoice == "scissors" || userChoice == "scissors" && opponentChoice == "rock" {
		fmt.Printf("Opponent chose %v! They win", opponentChoice)
		result = "Loss"

	}
	if userChoice == "rock" && opponentChoice == "scissors" || userChoice == "paper" && opponentChoice == "rock" || userChoice == "scissors" && opponentChoice == "paper" {
		fmt.Printf("Opponent chose %v! You win", opponentChoice)
		result = "Win"
	}
	if userChoice == "rock" && opponentChoice == "rock" || userChoice == "paper" && opponentChoice == "paper" || userChoice == "scissors" && opponentChoice == "scissors" {
		fmt.Printf("Opponent chose %v! Draw", opponentChoice)
		result = "Draw"
	}
	return result

}

func addScore(userScore int, opponentScore int, result string) (int, int) {
	if result == "Win" {
		userScore++
	}
	if result == "Loss" {
		opponentScore++
	}
	return userScore, opponentScore
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
