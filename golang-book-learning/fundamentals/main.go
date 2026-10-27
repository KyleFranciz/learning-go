package main

import (
	"fmt"

	"golang-book-learning/fundamentals/consts"
	"golang-book-learning/fundamentals/copy"
	"golang-book-learning/fundamentals/make"
	mapsex "golang-book-learning/fundamentals/mapsEx"
	"golang-book-learning/fundamentals/moreaboutslices"
	"golang-book-learning/fundamentals/stringsandrunes"

	compositetypes "golang-book-learning/fundamentals/composite_types"
)

func main() {
	// Declared num var in function. Use walrus to declar and type new vars
	exampleNum := 26 // NOTE: walrus declaration is not ideal when making variables for packages to use. They have to be inside a function.
	exampleNum += 2
	// var ex int          // this declares a variable and set it to 0
	// x, y := 10, "money" // another way to declare variables and types are infered

	const username = "KyleDev" // this is just the name of a value, constants are essentially taggin a value
	username2 := "DevKyle"     // vars have a memory address, they can be modified and changed. Constants can't, they do not have a memory address.

	if exampleBool {
		// doesn't run since condition is false
		fmt.Printf("test: %d", exampleNum)
	}
	if secondBoolExample {
		// runs because condition is set to true
		fmt.Printf("test 2: %s", name)
		fmt.Println(" ")
		ComplexExample(12.0, 7.3)
		ConversionExample(2, 12.5)
		fmt.Printf("Username: %s, Other name: %s ", username, username2)
		// examples from the const package to show the different types of constants
		fmt.Println(consts.UntypedConstExample())
		fmt.Println(consts.TypedConstExample())

		// array test
		fmt.Printf("%#v", compositetypes.TestArrayOuput())
		fmt.Printf("%#v", compositetypes.SparseArrayExample())

		// test slice
		testSlice := []int{1, 2, 3, 4, 5}

		// example of merging two slices into each other
		fmt.Printf("%#v", moreaboutslices.MergeSliceExample(testSlice))

		moreaboutslices.CapacityUnderstanding()

		dummySlice := make.MakeSliceExample(12) // custom function helps w/ making a slice w custom sizing to show slices are

		fmt.Printf("%#v", dummySlice)

		fmt.Println(" ")

		fmt.Println(cap(make.TestingMakeCapacity()))

		newSlice := append(make.TestingMakeCapacity2(), 1, 2, 3, 4, 5) // [1,2,3] - make function made 2 slot get filed and capacity is set to 5
		fmt.Println(cap(newSlice))
		fmt.Println(newSlice)

		exampleSlice := []int{1, 2, 3, 4, 5}

		fmt.Println(moreaboutslices.SliceSclices(exampleSlice))

		ProperSlice := append(moreaboutslices.ProperSubsliceAppend(), 2)
		fmt.Println("Appending Slice Proper Way:", ProperSlice)
		fmt.Println("Appending Slice Proper Way Cap: ", cap(ProperSlice)) // 2
		ProperSlice2 := append(moreaboutslices.ProperSubsliceAppend(), 4, 2, 4)
		fmt.Println("Appending Slice Proper Way (+2):", ProperSlice2)
		fmt.Println("Appending Slice Proper Way Cap: ", cap(ProperSlice2)) // 6

		// NOTE: Slice that is slice have two different memeory address, "x" doesnt affect the base subslice at all. 20 is only added to x
		x := moreaboutslices.ProperSubsliceAppend()
		x = append(x, 20)
		fmt.Println("Appending 20 to tripple sliced slice: ", x)
		fmt.Println("Base slice: ", moreaboutslices.ProperSubsliceAppend())

		fmt.Println("Copy ended up cloning (y) into (x) and this is what (x) now has:", copy.ExampleCopy())

		fmt.Println("This is what was copied from (y) over to (x): ", copy.AnotherCopyExample())

		fmt.Println("This function is the more modern way to copy a slice into a new one slices.Clone(slice) ")
		fmt.Println("This is the copy that was made:", copy.ModernCopyExample())
		fmt.Println("This is the length of the copy that was made:", len(copy.ModernCopyExample()))

		fmt.Println(stringsandrunes.SlicesAndRuneExample())
		fmt.Println("This is what a map looks like (map[string][]string): ", mapsex.MapUnderstanding())
		fmt.Println("This is the len of a map made using the make function: ", len(mapsex.MapExampleUsingMake())) // NOTE: len only tells the amount of key value pairs when used on a map!

		// maps: finding if key exists in map
		fmt.Println("This is an example of using the ok idiom with the map to check if something exists in it")
		v, ok := mapsex.MapOkExample("Pie")
		fmt.Println("This is how map return values works when trying to do look ups ", v, ok) // doesn't return the value of the entered key, and returns false
		v2, ok2 := mapsex.MapOkExample("Number")
		fmt.Println("This is how map return values works when trying to do look ups ", v2, ok2) // returns the value of the entered key, and returns true

		// example of changing vals in maps, enter names: Kyle, Karim, Cam in to see the different incrementations
		fmt.Println("This is an example of a key's value being incremented", mapsex.IncreasingMapValueExample("Cam"))

	}
}

// this is a var example of multiple vars
var (
	// multiple variables can be made like this at one time. Doesnt have to be the same type
	exampleBool       bool     // sets to false by default
	secondBoolExample = true   // this is defining a bool
	name              = "Kyle" // type is infered automatically when declared and set
)

func ComplexExample(x float32, y float32) {
	// this function prints out an imaginary number
	num1 := complex(x, y)
	num2 := complex(x, y)
	fmt.Println(num1 + num2)
}

func SmallExample() {
	fmt.Println("This is an example, othre real function till be added hare")
}

func ConversionExample(x int, y float32) {
	// example of conversion
	floatExample := float32(x) + y // convert int  to a float
	intExample := int(y) + x       // convert float into an int
	fmt.Println(floatExample, intExample)
}
