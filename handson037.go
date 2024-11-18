package main

import (
	"fmt"
)

func main() {

	m := map[string]int{"James": 42, "Moneypenny": 32}
	age := m["James"]
	fmt.Println(age)

	value, ok := m["Q"]

	if ok {
		fmt.Printf("Found Q with value: %v", value)
	} else {
		fmt.Printf("Q not found")
	}

}
