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
	found := false
	remainingGuesses := 10
	for remainingGuesses >= 1 {

		guessWordOrLetter := askUser(" Press 1 to guess a letter. Press 2 to guess the word")

		if guessWordOrLetter == "1" {

			userLetterGuess := askUser("Guess a letter in the word")
			isCorrect, hiddenWord := checkUserLetterGuess(chosenWordSplit, userLetterGuess, hiddenWord)
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
		if remainingGuesses == 0 {
			handleGameResult(remainingGuesses, chosenWord, found)
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
