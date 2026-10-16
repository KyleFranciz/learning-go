package main // dsa questions

// check for dupes
func hasDuplicates(nums []int) bool {
	// create the set manually with the make func and the map constructor
	seen := make(map[int]bool)

	// for range in number (no for in equivalent like in java and python)(idx and value used for range, have to handle it _ (makes it so that it doesn't get checked))
	for _, number := range nums {
		// check in the set using the value if the number exists in there
		if seen[number] {
			// return true if it exists in made set
			return true
		}
		// the value in the set is made true, the number stored in the set is the key
		seen[number] = true
	}
	// return false in the end if the number is not stored in the set
	return false
}

// main function to run the example
func main() {
	// example
	example := []int{1, 2, 3, 5}
	// show the example
	println(hasDuplicates(example))
}
