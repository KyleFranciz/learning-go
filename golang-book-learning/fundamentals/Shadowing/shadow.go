// Package shadowing: just a package to help understand how if works
// NOTE: as well that I have a shadow vet package for go installed now, can add to a make file to help when checking
package shadowing

import "fmt"

func UnderstandingShadows1() {
	x, y := 10, 20

	if x > 5 {
		// prints the first outer scope declaration
		fmt.Println(x, y)
		// assigns shadow variable w/in this block
		x := 5
		// prints shadow x
		fmt.Println(x)
	}
	// prints outer x because were back in the outer scope
	fmt.Println(x)
}

func UnderstandingShadows2() {
	// universal block variable true is bool the first time
	fmt.Println(true)
	// universal block true is set to "t"
	true := "t"
	// true is now "t" if printed
	fmt.Println(true)
}
