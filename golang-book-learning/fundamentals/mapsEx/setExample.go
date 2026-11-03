// Package mapsex
package mapsex

import (
	"fmt"
)

func SetExample() {
	setEx := map[int]bool{}                      // dummy set to store the key and value as true
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // slice of numbers to add to the set

	// for loop to add in the values to the set
	// index and value are the return values
	for _, v := range vals {
		// add the int keys and give them the value of true since the int is now in the set
		setEx[v] = true
		// cneck if the v(key) that are in the vals array are greater than 5
		if v > 5 {
			fmt.Println("numbers greater than 5 are in the set")
			if v < 5 {
				fmt.Println(setEx[7]) // returns bool if the value is in the set
				continue              // break also works the same
			}
		}
	}
}

func SetUsingStruct() {
	intSet := map[int]struct{}{}
	values := []int{2, 4, 5, 6, 7, 10, 11}

	for _, v := range values {
		// NOTE: type is struct{}, this is then with no value = {}
		intSet[v] = struct{}{} // less declarative but more maluable based on the data type
		// matters for memory management
	}
	if _, ok := intSet[5]; ok {
		fmt.Println("5 is a good number")
	}
}
