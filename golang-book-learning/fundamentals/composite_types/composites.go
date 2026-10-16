// Package compositetypes: is to show how arrays, maps, slices, and array like obj work in go
package compositetypes

func SliceExample() []int {
	// returns an empty slice
	var example []int
	return example
}

func ArrayExample2() []int {
	example2 := [3]int{2, 4, 6} // array size of 3 initialized and variable set for each position
	slice := example2[:]        // this creates a copy of the array and creates a slice, the array still exists in memory but we get a slice w/ the same numbers as the array
	return slice
}

func SparseArrayExample() [5]int {
	sparse := [5]int{1, 2: 3, 4: 5} // numbers like 2: and 5: will be set to their zero value (0) in he array
	return sparse
}

func TestArrayOuput() [5]int {
	arrayTest := [5]int{1, 3, 4, 6, 7}
	return arrayTest
}

func ArrayLiteralExample() [3]int {
	x := [...]int{1, 2, 3} // this automatically counts that amount of numbers in the array, still have to adjust the number in the return type
	return x
}
