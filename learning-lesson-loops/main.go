package main

import (
	"fmt"
)

func main() {
	intArr := [...]int32{1, 3, 5, 7}
	firstMap := map[string]uint8{"Kyle": 26, "Omar": 24, "Mom": 61}

	// loop throught the map (dict)
	for name, age := range firstMap {
		// prints the name and the age
		fmt.Println(name, age)
	}

	// loop through the array
	for idx, value := range intArr {
		fmt.Println(idx, value) // prints the idx and the value of the position
	}

	// while loop
	idx := 0
	for idx < 11 {
		fmt.Println(idx)
		idx = idx + 1
	}

	// while loop conditional
	for idx2 := 0; idx2 < 11; idx2++ {
		// declare , condition , increment
		fmt.Println(idx2)
	}
}
