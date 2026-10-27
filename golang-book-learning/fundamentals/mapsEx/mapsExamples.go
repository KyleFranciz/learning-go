// Package mapsex: just understanding how maps work in go
// map[keyType]valueType - this is how maps are created and made
package mapsex

func MapUnderstanding() map[string][]string {
	// NOTE: This is a nil map, the length and return type will be zero, you declare it using var
	// var nilMap map[string]int

	// NOTE: This is how you create an empty map literal, it has a length of 0 but we can read and write to it
	classGroups := map[string][]string{"Bears": {"Kyle", "Mark", "Dean"}, "Penguins": {"Paula", "Marg", "Paul"}, "Cats": {"Macy", "Dina", "Luke"}}
	return classGroups
}

func MapExampleUsingMake() map[int][]string {
	// key: int and value: []string
	exampleMap := make(map[int][]string, 10)
	return exampleMap
}

func MapOkExample(wordToCheckInMap string) (int, bool) {
	mapExmaple := map[string]int{
		"Number":      32,
		"OtherNumber": 2027,
	}

	v, alright := mapExmaple[wordToCheckInMap] // alright can be anything, usually is "ok" in best practices
	return v, alright
}

func IncreasingMapValueExample(name string) int {
	dummy := map[string]int{
		"Kyle":  0,
		"Karim": 0,
		"Cam":   0,
	}

	dummy["Kyle"] = 1 // value us set to 1
	dummy["Karim"]++  // increments that value if its and int
	dummy["Cam"] = 3  // value set to 3

	return dummy[name]
}
