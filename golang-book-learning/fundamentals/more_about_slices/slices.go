// Package moreaboutslices: This shows how slices work in go and the purpose for what they can be used for and how they work
package moreaboutslices

import "fmt"

func SliceCreatorExample() []int {
	// func makes a slice
	slice := []int{}
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
