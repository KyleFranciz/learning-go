// Package moreaboutslices: This shows how slices work in go and the purpose for what they can be used for and how they work
package moreaboutslices

import "fmt"

func SliceCreatorExample() []int {
	// func makes a slice
	slice := []int{} // returns a non-nil slice that has zero length
	// var slices []int // NOTE: other way to declare a slice
	return slice
}

func ModifySliceExample(slice []int, numberToAddIn int) []int {
	newSlice := append(slice, numberToAddIn)
	return newSlice
}

func GetLengthOfSliceExample(slice []int) int {
	// returns the len of the slice
	return len(slice)
}

func GetCapacityOfSlice(slice []int) int {
	// gets and returns the capacity of arrays, slices, maps, etc..
	return cap(slice)
}

func MergeSliceExample(slice []int) []int {
	mergeSlice := []int{0, 2}                 // created dummy slice that I can merge input slice into
	mergeSlice = append(mergeSlice, slice...) // add input slice into mergeSlice
	return mergeSlice
}

func CapacityUnderstanding() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))
}

func SliceSclices(slice []int) []int {
	return slice[1:] // NOTE: Appending to a subslice is bad usually because it affect the sublices because they share the same mem by default
}

func ProperSubsliceAppend() []int {
	// this is how if I wanted to add things to a slice, this is how it should be done
	example := make([]int, 0, 5)
	testSlice := append(example, 1, 2, 3, 4, 5)
	x := testSlice[2:4:4] // if the cap gets overflowed, the capacity get increased by the second number which is 2 in this case
	return x
}
