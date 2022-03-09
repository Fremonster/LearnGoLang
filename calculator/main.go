package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// "go build" to build it for windows (GOOS="windows" if not on windows machine)
// "GOOS="linux" go build" to build for linux

func main() {

	float1 := getInput("Value 1: ")
	float2 := getInput("Value 2: ")
	operation := getOperation()

	var total float64 = total(float1, float2, operation)
	fmt.Printf("The total of %v and %v with operation %v is %v\n\n", float1, float2, operation, total)

}

var reader = bufio.NewReader(os.Stdin)

func getInput(prompt string) float64 {
	fmt.Print(prompt)
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Print(err)
		//custom error message with panic
		panic("Value must be a number")
	}
	return float1
}

func getOperation() string {
	fmt.Print("Operation: ")
	operation, _ := reader.ReadString('\n')
	//:= use to assign value to a new variable, = used to reassign value to an existing variable
	operation = strings.TrimSpace(operation)
	if !strings.ContainsAny("+-*/", operation) {
		panic("Please enter a valid math operation: +-* or /")
	}
	return operation
}

// create 4 functions, one for each of +=*/
// one function that can handle input for value 1 and value 2
// should be able to receive value for operation also
// trim spaces of character and then use switch statement to determine operation

func add(value1 float64, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1 float64, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1 float64, value2 float64) float64 {
	return value1 * value2
}

func divide(value1 float64, value2 float64) float64 {
	return value1/value2
}

func total(value1 float64, value2 float64, operation string) float64 {
	var total float64
	switch operation {
	case "+": 
		total = add(value1, value2)
	case "-": 
		total = subtract(value1, value2)	
	case "*": 
		total = multiply(value1, value2)
	case "/": 
		total = divide(value1, value2)
	default:
		total = 0
		fmt.Println("Input was not a supported operation")	
	}
	total = math.Round(total*100)/100
	return total
}