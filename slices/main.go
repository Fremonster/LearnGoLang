package main

import (
	"fmt"
	"sort"
)

func main() {
	// an array
	var colorArr = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colorArr)
	// a slice sits on top of an array
	// runtime allocates the required memory and creates
	// the array in the background but returns the slice
	// all items in the slice are of the same type, but 
	// resizeable and can be sorted

	// a slice - remove number to create a slice instead of array
	var colorsSlc = []string{"Red", "Green", "Blue"}
	fmt.Println(colorsSlc)
	colorsSlc = append(colorsSlc, "Purple")
	fmt.Println(colorsSlc)

	// remove 0 index value - indicate a range 
	// start with index 1 and go to end of length of slice
	colorsSlc = append(colorsSlc[1:len(colorsSlc)])
	fmt.Println(colorsSlc)

	// start value of 0 is default for range
	// to remove last value in slice - deletes last item
	colorsSlc = append(colorsSlc[:len(colorsSlc) - 1])
	fmt.Println(colorsSlc)

	// using make - type, initial length, capacity that caps number of items slice can contain
	numbers := make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 32
	numbers[2] = 578
	numbers[3] = 89
	numbers[4] = 101
	fmt.Println(numbers)

	// can leave out capacity value, then can add more than 5 items
	numbers2 := make([]int, 5)
	numbers2[0] = 134
	numbers2[1] = 32
	numbers2[2] = 578
	numbers2[3] = 89
	numbers2[4] = 101
	numbers2 = append(numbers2, 657)
	fmt.Println(numbers2)

	// can sort slices
	sort.Ints(numbers2)
	fmt.Println(numbers2)
}