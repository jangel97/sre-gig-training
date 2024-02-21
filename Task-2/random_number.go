package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Function to generate a random number between 1 and 100
func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	return rand.Intn(100) + 1        // Return a random number between 1 and 100
}

// Function to evaluate and print messages based on the generated number
func evaluateAndPrintNumber(number int) {
	if number > 50 && number%2 == 0 {
		fmt.Println("It's closer to 100, and it's even!")
	} else if number > 50 {
		fmt.Println("It's closer to 100")
	} else if number == 50 {
		fmt.Println("It's 50!")
	} else {
		fmt.Println("It's closer to 0")
	}
	fmt.Println("Generated number:", number)
}

func main() {
	number := generateRandomNumber() // Generate a random number
	evaluateAndPrintNumber(number)   // Evaluate the number and print the appropriate message
}
