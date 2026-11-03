// Package structsex
package structsex

type Book struct {
	// can be scoped at any block level
	// NOTE: structs can be used to make types for data fetching, capital also determines private or public
	Title  string
	Author string
	Pages  int
	Year   int
}

func CreateBook(name string, author string, pages int, year int) *Book {
	// return returns the exact address of the book
	return &Book{
		name,
		author,
		pages,
		year,
	}
}

func (book *Book) Rename(newTitle string) {
	// return doesnt need to happen, the obj directly is being changed
	book.Title = newTitle
}

func (book *Book) UpdatePages(newPageNumber int) {
	// return doesnt need to happen, the obj directly is being changed
	book.Pages = newPageNumber
}

func (book *Book) UpdateAuthor(authorName string) {
	book.Author = authorName
}

func (book *Book) UpdateYearPublished(newYear int) {
	book.Year = newYear
}

func (book *Book) GetTitle() string {
	// return the title of the book (string)
	return book.Title
}

func (book *Book) GetAuthor() string {
	return book.Author
}

func (book *Book) GetYearPublish() int {
	return book.Year
}
