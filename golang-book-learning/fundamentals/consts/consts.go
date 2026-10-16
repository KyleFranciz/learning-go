// Package consts is just to explain the different ways that I can use constants in go
package consts

func UntypedConstExample() string {
	// this is what an untyped const looks like
	const name = "Mark" // regular const
	return name
}

func TypedConstExample() int {
	// this is what an typed const looks like
	const typedX int = 10 // typed const
	return typedX
}
