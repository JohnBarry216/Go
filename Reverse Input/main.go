package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/**
 * Takes a string input from the user and reverses
 */
func main() {

	var userInput string

	// Ask user to enter a string
	fmt.Println("Please enter a string: ")
	in := bufio.NewReader(os.Stdin)
	userInput, err := in.ReadString('\n')

	// simple error checking
	if err != nil {
		log.Fatalf("Error Occured: %s", err)
	}

	fmt.Print("\nYour reversed string is:")
	fmt.Println(reverseString(userInput))

}

/**
 * takes a string and reverses it
 * @return string
 */
func reverseString(s string) string {

	// convert stirng to an array
	var orgRunes []rune = []rune(s)
	var newRunes []rune = make([]rune, len(s))

	// loop for the length of the string
	for i := range s {
		newRunes[i] = orgRunes[len(s)-1-i]
	}

	return string(newRunes)
}
