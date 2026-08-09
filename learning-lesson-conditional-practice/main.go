package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// main function
func main() {
	// create score variables
	playerScore := 0
	computerScore := 0

	for {
		// create new scanner
		scanner := bufio.NewScanner(os.Stdin)

		greatUser()                          // greet user
		checkResult := getUserInput(scanner) // get the user input
		randomPick := getComputerPick()      // get the computer input

		playerScore, computerScore = checkIfWin(checkResult, randomPick, playerScore, computerScore)

		// display picks
		fmt.Println("You picked:", checkResult)
		fmt.Println("Computer picked:", randomPick)
		fmt.Println("Player score:", playerScore)
		fmt.Println("Computer score:", computerScore)

		verdict := keepPlaying(scanner)

		if !verdict {
			break
		}

		// other wise the loop is continued
	}
}

func keepPlaying(scanner *bufio.Scanner) bool {
	fmt.Println("Do you want to keep playing: y or n")

	for {
		scanner.Scan()
		answer := scanner.Text()
		switch answer {
		case "y":
			return true
		case "n":
			return false
		default:
			fmt.Println("Please enter y or n!")
		}
	}
}

// function to greet the user
func greatUser() {
	fmt.Println("Hey there player!!!")
	fmt.Print("Ready to get started?!!")
}

// function to get the users input
func getUserInput(scanner *bufio.Scanner) string {
	for {
		fmt.Println("Enter: rock (r), paper(p) or scissors (s)")
		// get the users input
		scanner.Scan()
		// store the user input
		userInput := scanner.Text()
		// make lower case in case
		userInput = strings.ToLower(userInput)
		// return the users input to be conditionally checked
		validatedChoice := validateChoice(userInput)

		// check to make sure the returned value isn't empty
		if validatedChoice != "" {
			return validatedChoice
		}

	}
}

// function to make sure all the choices match
func validateChoice(choice string) string {
	switch choice {
	case "r", "p", "s":
		return choice
	default:
		fmt.Println("Please enter: r , s or p (try again)")
		return ""

	}
}

// function to get the computer's choice
func getComputerPick() string {
	choices := []string{"r", "p", "s"}
	// randomPick
	randomIndex := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(choices))

	// return a random pick
	return choices[randomIndex]
}

// function to check which player wins
func checkIfWin(userInput string, randomPick string, playerScore int, computerScore int) (int, int) {
	// exit early if they cancel out
	if userInput == randomPick {
		return playerScore, computerScore
	}

	if userInput == "r" && randomPick == "s" ||
		userInput == "p" && randomPick == "r" ||
		userInput == "s" && randomPick == "p" {
		playerScore += 1
	} else {
		computerScore += 1
	}

	return playerScore, computerScore
}
