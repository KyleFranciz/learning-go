package main // <- special name to help with compling the full project

// imports
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv" // handles string conversion
	"strings" // handles checking for space
)

func main() {
	dummyNum := 12.5
	testingNumbers(dummyNum)
	answer, err := testingAddition()
	fmt.Println(answer, err)
	testingText("peter parker")

	result, err := errorHandlingNumber(5) // NOTE: this function can possibly return an error

	// NOTE: this is where we accout for the error that I get from the function
	if err != nil {
		fmt.Printf("There has beenn an error %v", err)
	} else {
		// otherwise we print the result
		fmt.Println(result)
	}
}

// learning error handling
func errorHandlingName(name string) (string, error) {
	// declare a error variable
	var err error

	// lower the name
	lowerName := strings.ToLower(name)

	// if the name != kyle raise an error
	if lowerName != "kyle" {
		err = errors.New("you are not kyle")
		return "", err
	} else {
		fmt.Println("You are a loser")
	}

	return "Hey Kyle", err
}

// error handling numbers
func errorHandlingNumber(number int) (int, error) {
	var err error

	if number != 12 {
		err = errors.New("the number is not the secret code")
		return 0, err
	}

	return number, err
}

// understanding variables
func variableUnderstanding() (string, int, int, string) {
	runeExample := "k"
	var1, var2 := 1, 2
	const name string = "Kyle" // NOTE: A constants variable can't be changed once declared and set
	fmt.Println(runeExample, var1, var2, name)
	return runeExample, var1, var2, name
}

func testingText(word string) {
	fmt.Println(word)
}

// this is an example of how functions work in go
func testingNumbers(dummyNumber float64) {
	fmt.Println("Here is a float number", dummyNumber)
}

func testingAddition() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the first number you want to add: ")
	// check when the user hits enter
	firstInput, _ := reader.ReadString('\n')
	firstInput = strings.TrimSpace(firstInput)
	// check when the user hits enter
	secondInput, _ := reader.ReadString('\n')
	secondInput = strings.TrimSpace(secondInput)

	// convert the string to number
	firstNumber, err := strconv.Atoi(firstInput)
	// error handling
	if err != nil {
		return 0, err
	}
	// convert the string to number
	secondNumber, err := strconv.Atoi(secondInput)
	// error handling
	if err != nil {
		return 0, err
	}
	answer := firstNumber + secondNumber
	return answer, nil
}
