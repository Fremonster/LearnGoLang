package main

import (
	"fmt"
)

const aConst int = 64

func main() {

	// explicit typing
	var aString string = "This is Go!!!"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", aString)

	var anInteger int = 42
	fmt.Println(anInteger)
	fmt.Printf("The variable's type is %T\n", anInteger)

	var defaultInt int // defaults to 0
	fmt.Println(defaultInt)

	// implicit typing
	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("The variable's type is %T\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variable's type is %T\n", anotherInt)

	// := operator - only works inside of functions
	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("The variable's type is %T\n", myString)

	// constants - defined outside of the function
	fmt.Println(aConst)
	fmt.Printf("The variable's type is %T\n", aConst)

	// how to create a map - memory allocation
	// new() allocates but does not initialize memory
	// make() allocates and initializes memory
	m := make(map[string]int)
	m["theAnswer"] = 42
	fmt.Println(m)

}
