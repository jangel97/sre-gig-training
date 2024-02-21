package main

import "fmt"

// Function for Task 1: Manage menu items
func manageMenu() {
	// Define a slice for the menu
	menu := []string{}

	// Append items
	menu = append(menu, "hamburger")
	menu = append(menu, "salad")

	// Iterate and print
	for _, item := range menu {
		fmt.Printf("Food: %s\n", item)
	}
}

// Function for Task 2: Display items and their indexes
func displayItems() {
	// Define an array of 5 items
	items := [5]string{"Item1", "Item2", "Item3", "Item4", "Item5"}

	// Iterate and print each item with its index
	for index, item := range items {
		fmt.Printf("This is %s and its index in the array is %d\n", item, index)
	}
}

func main() {
	// Call the function for Task 1
	manageMenu()

	// Call the function for Task 2
	displayItems()
}
