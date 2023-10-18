package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

//Hangman tasks
//Choose a random word(from a given list or a csv file)
//split each letter into slices
//create a user guess input
//check if letter is in chosen word

func main() {

	chosenWord := getWord()
	letterArray := strings.Split(chosenWord, "")

	fmt.Print("Your word is ")
	hiddenWord := printLettersInWord(chosenWord)
	fmt.Println(hiddenWord)

	playGame(letterArray, hiddenWord, chosenWord)

}

func playGame(letterArray []string, hiddenWord []string, chosenWord string) {
	found := false
	remainingGuesses := 10
	for remainingGuesses >= 1 {

		guessWordOrLetter := askUser(" Press 1 to guess a letter. Press 2 to guess the word")

		if guessWordOrLetter == "1" {

			userLetterGuess := askUser("Guess a letter in the word")
			isCorrect, hiddenWord := checkUserLetterGuess(letterArray, userLetterGuess, hiddenWord)
			found = checkHiddenWord(hiddenWord, chosenWord, found)

			if found {
				remainingGuesses = 0
			}

			if isCorrect == "incorrect" {
				remainingGuesses--

			}

			fmt.Printf("The letter is %v. You have %v guesses left!", isCorrect, remainingGuesses)
			fmt.Println("")
			fmt.Println("Your word is", hiddenWord)

		} else if guessWordOrLetter == "2" {

			userWordGuess := askUser("Guess the word")
			found = checkUserWordGuess(userWordGuess, chosenWord, found)
			remainingGuesses = 0

		} else {
			fmt.Println("Enter a valid input")
		}
		if remainingGuesses == 0 && !found {
			fmt.Println("You lost! The word was", chosenWord)
		}
		if remainingGuesses == 0 && found {
			fmt.Println("Congrats! You Win!")
		}
	}

}

func getWord() string {
	words := []string{"child", "dog", "animal", "help"}
	arrayLength := len(words)
	indexNumber := rand.Intn(arrayLength)
	chosenWord := words[indexNumber]

	return chosenWord
}

func printLettersInWord(chosenWord string) []string {
	hiddenWord := make([]string, len(chosenWord))
	for i := range chosenWord {
		hiddenWord[i] = "_"
	}
	return hiddenWord
}

func checkUserLetterGuess(letterArray []string, userLetterGuess string, hiddenWord []string) (string, []string) {
	isCorrect := "incorrect"
	for i, letter := range letterArray {

		if userLetterGuess == letter {
			hiddenWord[i] = userLetterGuess
			isCorrect = "correct"
		}

	}
	return isCorrect, hiddenWord

}

func checkHiddenWord(hiddenWord []string, chosenWord string, found bool) bool {
	hiddenWordJoined := strings.Join(hiddenWord, "")
	if hiddenWordJoined == chosenWord {
		found = true
	}
	return found
}

func checkUserWordGuess(userWordGuess string, chosenWord string, found bool) bool {
	if userWordGuess == chosenWord {
		fmt.Println("You guessed the word! Well done")
		found = true

	} else {
		fmt.Printf("Unlucky, your guess is incorrect.")

	}
	return found

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
