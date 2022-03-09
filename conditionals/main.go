package main

import (
	"fmt"
)

func main() {
	
	theAnswer := 42
	var result string

	if theAnswer < 0 {
		result = "Less than zero"
	} else if theAnswer == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)

	// can include an initial statement that is part of if statement
	// for example, can initialize variables, like x := -42
	if x := -42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)
}