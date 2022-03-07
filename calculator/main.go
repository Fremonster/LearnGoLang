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
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		//custom error message with panic
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	//:= use to assign value to a new variable, = used to reassign value to an existing variable
	sum := float1 + float2
	sum = math.Round(sum*100)/100
	fmt.Printf("The sum of %v and %v is %v\n\n", float1, float2, sum)

}