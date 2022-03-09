package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// can use range keyword
	for i := range colors {
		fmt.Println(colors[i])
	}

	// like a for each statment
	// for index, value - if don't need index, put underscore
	for _, color := range colors {
		fmt.Println(color)
	}

	// like a while statement
	value := 1
	for value < 10 {
		fmt.Println("Value", value)
		value++
	}

	// break, continue, goto
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum", sum)
		if sum == 42 {
			fmt.Println("42 is the answer")
			break
		}
		if sum < 500 {
			continue
		}
		if sum > 200 {
			goto theEnd
		}
	}

	theEnd : fmt.Println("End of program")

	var out int
	for j:=0; j < 20; j++ {
		out=j*j + out
		if out > 12 {goto theVeryEnd }  }
	theVeryEnd: fmt.Println(out)  

}