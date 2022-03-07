package main

import (
	"fmt"
)

func main() {
	// astrisk means that it's a pointer, defaults to nil
	var p *int
	fmt.Println("Value of p is", p)
	//fmt.Println("Value of p", *p) - this would cause an error

	anInt := 42
	// & means pointing at memory address of the variable, not it's value
	var p2 = &anInt
	fmt.Println("Value of p2 is", *p2)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value1: ", *pointer1)

	*pointer1 = *pointer1 * 31
	fmt.Println("Pointer1: ", *pointer1)
	//changing value of the pointer1 also changes the value of value1 - same memory address
	//similar to reference variables in Java
	//even more similar to pointers in c
	fmt.Println("Value1: ", value1)

}
