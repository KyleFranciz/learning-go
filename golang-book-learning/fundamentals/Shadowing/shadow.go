// Package shadowing: just a package to help understand how if works
// NOTE: as well that I have a shadow vet package for go installed now, can add to a make file to help when checking
package shadowing

import "fmt"

func UnderstandingShadows1() {
	// set the variables
	x, y := 10, 20

	if x > 5 {
		// prints the first outer scope declaration
		fmt.Println(x, y)
		// assigns shadow variable w/in this block
		x, y := 5, 2
		// prints shadow x
		fmt.Println(x, y)
	}
	// prints outer x because were back in the outer scope
	fmt.Println(x, y) // prints 10
}

func UnderstandingShadows2() {
	// NOTE: universal block variable "true" is bool the first time
	fmt.Println(true)
	// universal block var true is set to "t"
	true := "t"
	// true is now "t" if printed out going forward, no longer a bool
	fmt.Println(true)
}
