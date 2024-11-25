package main

import "fmt"

func main() {
	/*
		Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
		"James", "Bond", "Shaken, not stirred" "Miss",
		"Moneypenny", "I'm 008."
		Range over the records, then range over the data in each record.
	*/

	slice1 := []string{"James", "Bond", "Shaken, not stirred"}

	slice2 := []string{"Miss", "Moneypenny", "I'm 008."}

	slice3 := [][]string{slice1, slice2}

	for _, v := range slice3 {
		fmt.Println(v)
		for a, b := range v {
			fmt.Println(a, b)

		}
	}
}
