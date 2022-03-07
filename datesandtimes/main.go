package main

import (
	"fmt"
	"time"

)


func main() {

	n := time.Now()
	fmt.Println("It is now", n)

	t := time.Date(2009, time.November, 10, 23, 0, 0 , 0, time.UTC)
	fmt.Println("Go launched at", t)
	fmt.Println(t.Format(time.ANSIC))

	// error that needs to be ignored with underscore or won't compile
	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)

	panic("This is my error message")

}