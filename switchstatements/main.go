package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	// get a number between 1 an 7
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	// don't need break statements
	var result string
	switch dow { // could include statement above if want to: switch dow := rand.Intn(7) + 1; h dow {	
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Monday"		
	case 3:
		result = "It's Tuesday"
	case 4:
		result = "It's Wednesday"
	case 5:
		result = "It's Thursday"
	case 6:
		result = "It's Friday"
	default:
		result = "It's Saturday"
	}
	fmt.Println(result)

	// could include statement in switch statement if want to
	// can also use "fallflow" to use flow similar to c and java
	// where falls through to next case unless use break statement
	switch dow -= 1; dow {	
	case 1:
		result = "It's Sunday"
		fallthrough
	case 2:
		result = "It's Monday"
		fallthrough	
	case 3:
		result = "It's Tuesday"
		fallthrough
	case 4:
		result = "It's Wednesday"
		fallthrough
	case 5:
		result = "It's Thursday"
		fallthrough
	case 6:
		result = "It's Friday"
		fallthrough
	default:
		result = "It's Saturday"
	}
	fmt.Println(result)
}