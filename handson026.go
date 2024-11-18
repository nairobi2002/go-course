package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("Value x: %v \n", x)
	fmt.Printf("Value y: %v \n", y)

	/*if x < 4 && y < 4 {
		fmt.Printf("Value x: %v, and Value y: %v, are less than 4", x, y)
	} else if x > 6 && y > 6 {
		fmt.Printf("Value x: %v, and Value y: %v, are greater than 6", x, y)
	} else if x >= 4 && x <= 6 {
		fmt.Printf("Value x: %v is between 4 and 6", x)
	} else if y != 5 {
		fmt.Printf("Value y: %v is not equal to 5", y)
	} else {
		fmt.Printf("None of the above")
	}*/

	switch {
	case x < 4 && y < 4:
		fmt.Printf("Value x: %v, and Value y: %v, are less than 4", x, y)
	case x > 6 && y > 6:
		fmt.Printf("Value x: %v, and Value y: %v, are greater than 6", x, y)
	case x >= 4 && x <= 6:
		fmt.Printf("Value x: %v is between 4 and 6", x)
	case y != 5:
		fmt.Printf("Value y: %v is not equal to 5", y)
	default:
		fmt.Printf("None of the above")
	}

}
