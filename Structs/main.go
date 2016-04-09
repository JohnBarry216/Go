package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

/**
 * Takes a string input from user and parses it for 2 numbers
 * assigns those to height and width respectivly and outputs calculated area
 */
func main() {

	rect1 := new(rect)

	// Ask user to enter rectangle values
	fmt.Println("Please enter the height and width of the rectangle you want to create (ex: 12 by 20)")

	// get user input
	in := bufio.NewReader(os.Stdin)
	userInput, err := in.ReadString('\n')

	// error check user input
	if err != nil {
		fmt.Printf("Error occured: %v", err)
		os.Exit(1)
	}

	// get user entered width and height and set rect properties
	userWidth, userHeight := parseUserInput(userInput)
	rect1.width = userWidth
	rect1.height = userHeight

	// print area of rectangle
	fmt.Println(rect1.getArea())
}

/**
 * parses strings for numbers
 */
func parseUserInput(userInput string) (int, int) {
	const NUMBERS_TO_FIND int = 2

	// regex to find numbers in string
	reg := regexp.MustCompile("[0-9]+")

	// get numbers. NOTE: the 2 means we only want 2
	var userNumbers []string = reg.FindAllString(userInput, NUMBERS_TO_FIND)

	// check if we have valid amout of numbers
	if userNumbers == nil || len(userNumbers) != NUMBERS_TO_FIND {
		fmt.Println("You must enter two numbers!")
		os.Exit(0)
	}

	// regex confirms that numbers are found, no need to error check here
	userHeight, _ := strconv.Atoi(userNumbers[0])
	userWidth, _ := strconv.Atoi(userNumbers[1])

	return userHeight, userWidth
}

/**
 * define rect struct
 */
type rect struct {
	width  int
	height int
}

/**
 * get area of related rect
 * @return int
 */
func (r rect) getArea() int {
	return r.width * r.height
}

/**
 * Sets the width of the related rect
 * @param  width int
 */
func (r *rect) setWidth(width int) {
	r.width = width
}

/**
 * Sets the height of the related rect
 * @param  height int
 */
func (r *rect) setHeight(height int) {
	r.height = height
}
