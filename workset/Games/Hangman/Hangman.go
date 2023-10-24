package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

//check if user has already guessed that letter
// make function for i in range user guesses, if guess equals appended array of user guesses, tell user to have another go

func main() {

	chosenWord := getWord()
	chosenWordSplit := strings.Split(chosenWord, "")

	hiddenWord := printLettersInWord(chosenWord)
	fmt.Println("Your word is", hiddenWord)

	playGame(chosenWordSplit, hiddenWord, chosenWord)

}

func playGame(chosenWordSplit []string, hiddenWord []string, chosenWord string) {
	var userGuesses []string
	isWordFound := false
	remainingGuesses := 10
	for remainingGuesses > 0 {

		guessWordOrLetter := askUser(" Press 1 to guess a letter. Press 2 to guess the word")

		if guessWordOrLetter == "1" {
			fmt.Println("The current letters you have guessed are:", userGuesses)
			userLetterGuess := askUser("Guess a letter in the word")

			duplicateGuess := checkForDuplicateGuess(userLetterGuess, userGuesses)

			if duplicateGuess {
				fmt.Println("You already guessed this letter, input another letter")
				continue
			}
			userGuesses = storeGuesses(userLetterGuess, userGuesses)

			isCorrect, hiddenWord := checkUserLetterGuess(chosenWordSplit, userLetterGuess, hiddenWord)
			isWordFound = checkHiddenWord(hiddenWord, chosenWord, isWordFound)

			if isWordFound {
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
			isWordFound = checkUserWordGuess(userWordGuess, chosenWord, isWordFound)
			remainingGuesses = 0

		} else {
			fmt.Println("Enter a valid input")
		}
		if remainingGuesses == 0 {
			handleGameResult(remainingGuesses, chosenWord, isWordFound)
		}
	}
}

func getWord() string {
	file, err := os.Open("hangmanWords.csv")

	if err != nil {
		log.Fatal("Error reading file!")
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records")
	}

	indexNumber := rand.Intn(len(records))
	chosenWord := records[indexNumber][0]

	return chosenWord
}

func printLettersInWord(chosenWord string) []string {
	hiddenWord := make([]string, len(chosenWord))
	for i := range chosenWord {
		hiddenWord[i] = "_"
	}
	return hiddenWord
}

func storeGuesses(userLetterGuess string, userGuesses []string) []string {
	userGuesses = append(userGuesses, userLetterGuess)
	return userGuesses
}

func checkForDuplicateGuess(userLetterGuess string, userGuesses []string) bool {
	duplicateGuess := false
	for i := range userGuesses {
		if userLetterGuess == userGuesses[i] {
			duplicateGuess = true
			return duplicateGuess

		}
	}
	return duplicateGuess
}

func checkUserLetterGuess(chosenWordSplit []string, userLetterGuess string, hiddenWord []string) (string, []string) {
	isCorrect := "incorrect"
	for i, letter := range chosenWordSplit {

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

func handleGameResult(remainingGuesses int, chosenWord string, found bool) {
	if !found {
		fmt.Println("You lost! The word was", chosenWord)
	} else {

		fmt.Println("Congrats! You Win!")

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

/*func getWordFromArray()string{
	words := []string{"child", "dog", "animal", "help"}
	arrayLength := len(words)
	indexNumber := rand.Intn(arrayLength)
	chosenWord := words[indexNumber]

	return chosenWord
}*/
