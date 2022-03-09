package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 10, "Woof"}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %+v\nWeight: %+v\n", poodle.Breed, poodle.Weight)

	poodle.Speak()
	// can change the public (exported fields) of an object
	poodle.Sound = "Arf!"
	poodle.Speak()
	poodle.SpeakThreeTimes()
	// if call function time, what happens
	// still only barks 3 times, because when pass Dog in as a receiver,
	// a copy is made of it. It's not a reference (would have to use pointers)
	poodle.SpeakThreeTimes()
}

// Dog is a struct
type Dog struct {
	Breed string
	Weight int
	Sound string

}

// In Go, a method is a member of a type
// Need to pass in a receiver (an identifier and the type)
// pass in a reference to a dog object

// use name of function and word "is" for comment above function

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// no method overrides
// methods can return values

// SpeakThreeTimes is how the dog speaks loudly
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}