package forloops

import "fmt"

func BreakAndContinue(nums int) {
	// int var's only allow one var for "for range" loops
	for idx := range nums {
		if nums%2 == 0 {
			fmt.Println("Divisible by 2: ", nums)
			nums--
			continue // continue - skips over the rest of the code in the loop and "continues" the loop
		}
		if nums%3 == 0 {
			fmt.Println("Divisible by 3: ", nums)
			nums--
			continue // continue - skips over the rest of the code in the loop and "continues" the loop

		}
		// NOTE: this part gets skipped w/ "continue" from above
		fmt.Println(idx, "Final Print: ", nums)
		// decrease the value of num on each itteration
		nums--

	}
}

func ForRangeLoops() {
	// int array vars (slices & maps as well and other collection types), allow for 2 vars to be used for the "for range" block
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// for loop side holds the variables that are needed, range side just unpacks the array or number if it were there
	for idx, val := range arr {
		// print out each index & value
		fmt.Println(idx, val)
	}
}

func ForRangeLoopWithoutIdx() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// NOTE: go requires you to have the all the variables needed for functions/keywords to be declared. '_' is a way to avoid using the variable while still satisfying go
	for _, value := range arr {
		fmt.Println(value)
	}
}

func ForRangeLoopWithoutValue() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// NOTE: go requires you to have the all the variables needed for functions/keywords to be declared.
	// The second var can just be left out and you can just use the idx for arrays or the keys for maps, go allows it.
	for idx := range arr {
		fmt.Println(idx)
	}
}
