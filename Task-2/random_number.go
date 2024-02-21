package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to generate a random number between 1 and 100 using a local random generator
func generateRandomNumber() int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(100) + 1
}

// Function to evaluate and print messages based on the generated number using a switch statement
func evaluateAndPrintNumber(number int) {
	switch {
	case number > 50 && number%2 == 0:
		fmt.Println("It's closer to 100, and it's even!")
	case number > 50:
		fmt.Println("It's closer to 100")
	case number == 50:
		fmt.Println("It's 50!")
	default:
		fmt.Println("It's closer to 0")
	}
	fmt.Println("Generated number:", number)
}

func main() {
	number := generateRandomNumber() // Generate a random number
	evaluateAndPrintNumber(number)   // Evaluate the number and print the appropriate message
}
