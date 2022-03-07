package main

import (
	"fmt"
	"math"
)

func main() {

	var aFloat = 3.14159
	var rounded = math.Round(aFloat)
	fmt.Printf("Original: %v, Rounded: %v\n", aFloat, rounded)

	var anInt int = 5
	var anotherFloat float64 = 42
	sum := float64(anInt) + anotherFloat
	fmt.Printf("Sum %v, Type: %T\n", sum, sum)

	// don't need to add space after "sum:", Println does it for you
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)
	// no spaces added for Print
	fmt.Print("Integer sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("\nFloat sum:", floatSum) // 164.89999999999998 precision problem with float
	
	// best way to get a more precise float value
	floatSum = math.Round(floatSum*100 / 100)
	fmt.Println("Float sum is now:", floatSum) 

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference %.2f\n", circumference)

	// output is 148.5
	val := 148.4789
	fmt.Printf("%.1f\n", val)
}
