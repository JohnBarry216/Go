package main

import (
	"fmt"
)

/**
 * Test of interfaces using basic animal example
 */
func main() {

	// declare cat and dog
	dog1 := Dog{"Fido"}
	cat1 := Cat{"Snowball"}

	// run them through the talk function, this will call each structs defined speak function
	talk(dog1)
	talk(cat1)
}

// Calls the speak function of the related animal
func talk(a Animal) {
	fmt.Println(a.speak())
}

////////////////
// Interfaces //
////////////////
type Animal interface {
	speak() string
}

/////////////
// Structs //
/////////////
type Dog struct {
	name string
}

type Cat struct {
	name string
}

/////////////
// Methods //
/////////////
func (d Dog) speak() string {
	return "Bark"
}
func (c Cat) speak() string {
	return "Meow"
}
