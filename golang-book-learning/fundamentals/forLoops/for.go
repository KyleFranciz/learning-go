// Package forloops - this is a file that shows how for loops work in go
package forloops

import "fmt"

func ForLoopC() {
	// NOTE: For Loops in go are like for loops in C and java - this is the older way,
	for itterator := 0; itterator < 10; itterator++ {
		// prints the itterator out 9 times
		// NOTE: ":=" symbol has to be used in the for loop when initializing a var
		fmt.Println(itterator)
	}
}

func ConditionalOnlyFor() {
	// NOTE: Conditional Looping - only use for loops
	itterator := 1 // itterator has to be declared first
	// write the condition statement and it loops for that long
	for itterator < 20 {
		fmt.Println(itterator)
		itterator = itterator * 2
	}
}

func InfininiteLoop() {
	// NOTE: This loop will go on indefinitly
	for {
		fmt.Println("loop")
	}
}
