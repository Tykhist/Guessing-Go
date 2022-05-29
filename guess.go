package main

import (
	"fmt"
)

func ask() int {
  // Initiates guessing game by asking the user to guess a number
	var input int
	fmt.Println("I am thinking of a number between 1-100. What is it?")
	fmt.Println("Type -1 to quit.")
	
	// Get the input from the user
	fmt.Scan(&input)
	return input
}

func main() {
  //Currently the answer is 56, but in the near future I want to randomize it using my Quick_RNG_GO program
	var guess int

	// Indefinite loop containing the ask() function
	guess = ask()
	for guess != 56 {
		if guess == -1 {
			fmt.Println("Better luck next time...")
			break
		} else {
			fmt.Println("Nope, not", guess)
			guess = ask()
		}
	}

	if guess == 56 {
		fmt.Println("You are correct! You've won!!")
	}
}
