package main

import "fmt"

func main() {
	// makes a default array of [0,0,0]
	var intArr [3]int32      // declare the amount of spots in the array
	intArr[1] = 11           // change the second number in the array
	fmt.Println(intArr[0])   // first number in the array
	fmt.Println(intArr[1:3]) // prints numbers 1 and 2

	// "&" handles printing out the memory addresses
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// other way to initialize and array
	arrTwo := [5]int32{1, 2, 3, 4, 5}      // standard way to create an array
	arrThree := [6]int32{1, 2, 3, 4, 5, 6} // other way to create an array
	finalArr := [...]int32{1, 3, 6, 9}     // the amount of numbers in the array gets infered by however many in the array
	fmt.Println(arrTwo)
	fmt.Println(arrThree)
	fmt.Println(finalArr)

	// slices - essentially a brand new array that gets created with more capacity to add more numbers
	sliceExample := []int32{1, 2, 3}
	fmt.Println(sliceExample)
	sliceExample = append(sliceExample, 8) // this function adds 8 to the end of the new array that gets made
	sliceExample2 := []int32{8, 9}
	sliceExample = append(sliceExample, sliceExample2...) // adds the new dummy array to the previous array
	fmt.Println(sliceExample)

	// creating a slice
	sliceExample3 := make([]int32, 5, 10) // size of 10 and capacity of 10
	// empty slice
	fmt.Println(sliceExample3)

	// map - key value pair, a dict
	myMap := make(map[string]uint8)
	fmt.Println(myMap)

	// creating an actual map
	myMap2 := map[string]uint8{"Kyle": 26, "Random": 28}
	fmt.Println(myMap2)

	// can also check if a key is in the map
	age, ok := myMap2["Kyle"]

	if ok {
		fmt.Printf("The age is v%", age)
	} else {
		fmt.Println("The key is not in this map")
	}
}
