// Package structs
package structs

// can be scoped at any block level
type book struct {
	name   string
	author string
	pages  int
	year   int
}

// make the variable of type book
var brigeBook book


BookExample := book{
	name: "Atomic Habits",
	author:"Kyle",
	pages: 210,
	year: 2026,
}



