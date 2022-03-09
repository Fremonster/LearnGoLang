package main

import (
	"fmt"
)

func main() {
	// set properties in braces in same order they are declared in the struct
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)

	// to see all fields and their values
	fmt.Printf("%+v\n", poodle)
	// to access field values - use dot operator
	fmt.Printf("Breed: %+v\nWeight: %+v\n", poodle.Breed, poodle.Weight)
	// can also use dot operator to change a field value
	poodle.Weight = 9
	fmt.Printf("%+v\n", poodle)

}

// (control / - creates "//"" - shortcut to single line comment)
// structs are similar to java classes, but no inheritance
// Dog is a struct 
type Dog struct {
	// uppercase = public (exported), lowercase = private 
	// (stuct can be lowercase - dog, then can't created one)
	Breed string
	Weight int
}