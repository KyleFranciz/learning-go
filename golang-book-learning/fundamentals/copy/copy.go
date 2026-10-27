// Package copy: the main purpose of this function is to help w/ copying one slice into another, there is no regular way to do this.
// try any other way and the source slice is just having a pointer set to the original slice and only that gets affected
// a = b , then you try to make a change to "a" only "b" is changed
package copy

import (
	"fmt"
	"slices"
)

func ExampleCopy() []int {
	x := make([]int, 4)    // x is the destination in this case and needs to have a lenghth big enough to copy everything
	y := []int{1, 2, 3, 5} // this is the source, a created slice that I want to have cloned

	fmt.Println("This is the original slice (y): ", y)
	fmt.Println("The copy func returns the amount of elements coppied into the new slice which is: ", copy(x, y)) // copies everything from y into x
	return x                                                                                                      // copy returns the lenghth of amount of elements actually copied
}

func AnotherCopyExample() []int {
	x := make([]int, 2)
	y := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("This is what (y)'s slice has inside: ", y)

	copyResult := copy(x, y)                                                    // only 2 elements will be copied because "x" only has a lenghth of 2
	fmt.Println("Amount of elements copied by the copy function: ", copyResult) // returns 2, it returns the amount of elements actually copied

	return x
}

func ModernCopyExample() []int {
	y := []int{23, 46, 1, 12, 332, 56, 8}

	x := slices.Clone(y) // this method makes and returns a new slice w/ needed cap
	// NOTE: Copy gets made into a new array, all elements are carried over, the capacity is just not set
	// regular copy method will allow for you to use "make" to set the capacity and the lenghth
	return x
}
