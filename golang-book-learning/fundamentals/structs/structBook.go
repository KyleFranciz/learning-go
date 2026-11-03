// Package structs
package structs

type Book struct {
	// can be scoped at any block level
	// NOTE: structs can be used to make types for data fetching, capital also determines private or public
	Name   string
	Author string
	Pages  int
	Year   int
}

func NewBook(name string, author string, pages int, year int) *Book {
	// make the variable of type book
	return &Book{
		name,
		author,
		pages,
		year,
	}
}

func (book *Book) Rename(newName string) {
	// return doesnt need to happen, the obj directly is being changed
	book.Name = newName
}

func (book *Book) UpdatePages(newPageNumber int) {
	// return doesnt need to happen, the obj directly is being changed
	book.Pages = newPageNumber
}
