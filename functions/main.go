package main

import (
	"fmt"
)

func main() {
	doSomething()
	sum := addValues(5, 8)
	fmt.Println(sum)
	sum = addValues2(11, 12)
	fmt.Println(sum)
	total := addAllValues(1, 2, 3, 4, 5)
	fmt.Println(total)
	multicount, multilength := addAllValues2(3, 4, 6, 8, 11)
	fmt.Println(multicount)
	fmt.Println(multilength)

}

func doSomething() {
	fmt.Println("Doing something")
}

// return type goes after parentheses
func addValues(value1 int, value2 int) int {
	return value1 + value2
}

// can also leave out the type of one of the parameters
func addValues2(value1, value2 int) int {
	return value1 + value2
}

// can have a function that accepts an arbitrary number of variables
func addAllValues(values ...int) int {
	// parameters are an array that can be looped through
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

// can return more than one value from a function
// comma delimited list for return types
func addAllValues2(values ...int) (int, int) {
	// parameters are an array that can be looped through
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}