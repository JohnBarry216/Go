package main

import (
	"fmt"
	"os"
)

/**
 * This program creates a output.txt file with text "Hello World"
 */
func main() {

	// create file
	outputFile, err := os.Create("output.txt")

	// error check
	if err != nil {
		fmt.Printf("Error occured when creating file: \n %s", err)
		return
	}

	// Write text to file
	_, err = outputFile.WriteString("Hello World!")

	if err != nil {
		fmt.Printf("Error occured when writing to file: \n %s", err)
		return
	}

	fmt.Println("File successfully created and written to!")
}
