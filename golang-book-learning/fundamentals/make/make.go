// Package make: explains what the make function does
package make

func MakeSliceExample(size int) []int {
	// this function shows how I can make slices or arrays or what ever with length of the array or the capacity cusom made for the use case that we need it for.
	makeReturnValue := make([]int, size) // make function is used to create slices, maps, and arrays with custom sizes to house the values that are needed
	return makeReturnValue
}

func TestingMakeCapacity() []int {
	ex := make([]int, 5)
	ex = append(ex, 5) // NOTE: append always increases the length of a slice
	return ex
}

func TestingMakeCapacity2() []int {
	// NOTE: This makes a slice that has the length of 2 and the capacity of 5
	// the make function can be used to create slices with custom lengths and capacities for specific use cases
	// NOTE: Never have the cap less than the length of the array, it cases run time errors
	return make([]int, 0, 5) // able to hold up to 5 but has 0 length by default
}
