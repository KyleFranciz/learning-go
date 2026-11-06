// Package ifblocks - this file just shows examples of if blocks
package ifblocks

import (
	"fmt"
)

func IfExample() int {
	if age := 26; age == 26 {
		fmt.Println("Access Granted")
		return age
	} else if age == 18 {
		fmt.Println("Access Okay")
	} else {
		fmt.Println("You are not allowed")
	}

	return 0
}

func IfExample2(test int) {
	if n := test; n == 0 {
		fmt.Println("the number matched")
	}
}

func OutOfScope() {
	if n := 10; n == 0 {
		fmt.Println("This number is too low")
	} else if n > 5 {
		fmt.Println("This number is okay to use")
	}
	// fmt.Println(n) // the use of n would have been out of scope
}
