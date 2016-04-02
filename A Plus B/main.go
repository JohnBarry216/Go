package main

import (
	"fmt"
	"os"
)

/**
 * This Program will take take 2 inputs,
 * 	confirm they are integers -1000 to 1000
 * 	and print the sum of the two values
 */
func main() {

	// Declare vars
	var a int32
	var b int32

	// Ask and listen for user input
	fmt.Print("Please enter two integers betwen -1000 and 1000: ")
	_, err := fmt.Scan(&a, &b)

	// Run basic error check, exit if error occured
	if err != nil {
		fmt.Print("Error occured: ", err)
		os.Exit(0)
	}

	// Check to make sure values are between -1000 and 100
	if checkRanges(a, b) == false {
		fmt.Print("Both values must be between -1000 and 1000")
	} else {
		fmt.Print(a + b)
	}
}

/**
 * Checks if a and b are both between -1000 and 1000
 * @return bool
 */
func checkRanges(a, b int32) bool {

	if a >= -1000 && a <= 1000 && b >= -1000 && b <= 1000 {
		return true
	} else {
		return false
	}
}
