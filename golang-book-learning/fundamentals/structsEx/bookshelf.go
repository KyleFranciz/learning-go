// Package structsex
package structsex

type BookShelf struct {
	name  string
	books []Book
}

func CreateBookShelf(name string, books []Book) *BookShelf {
	return &BookShelf{
		// create an empty array of books
		name,
		books,
	}
}
