package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// seed the random struct
	rand.Seed(time.Now().UTC().UnixNano())

	// Create randomly generated slice of 250 ints
	var randInts []int = rand.Perm(int(rand.Int31n(1000)))

	for i := 0; i < 10; i++ {

		// find max value out of int slice
		var maxValue int = getMaxValueOfSlice(randInts)

		fmt.Println("The max value is: ")
		fmt.Println(maxValue)
	}

}

/**
 * Takes a []int and finds the maximum value
 * @para intSlice []int
 * @return int
 */
func getMaxValueOfSlice(intSlice []int) int {
	// init with whatever is first, this will likely be overwritten
	var maxValue int = intSlice[0]

	// loop through the slice
	// 	set highest value to 'maxValue'
	for i, _ := range intSlice {

		if intSlice[i] > maxValue {
			maxValue = intSlice[i]
		}
	}
	return maxValue
}
