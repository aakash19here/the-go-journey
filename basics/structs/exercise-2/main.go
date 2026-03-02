package main

import "fmt"

// 1. Define your struct here
type Book struct {
	ID         int
	Title      string
	Author     string
	IsBorrowed bool
}

func main() {
	// 2. Create your slice of books here
	library := []Book{
		{
			ID:         1,
			Title:      "Harry Potter",
			Author:     "J.K Rowling",
			IsBorrowed: false,
		},
		{
			ID:         2,
			Title:      "Do epic shit",
			Author:     "Ankur Warikoo",
			IsBorrowed: false,
		},
		{
			ID:         3,
			Title:      "48 Laws of Power",
			Author:     "Robert Greene",
			IsBorrowed: true,
		},
	}

	ToggleStatus(&library)

	for _, value := range library {
		if value.IsBorrowed {
			fmt.Println("The Book is", value.Title, "by", value.Author, "is borrowed", value.IsBorrowed)
		}
	}

}

func ToggleStatus(lib *[]Book) {
	for i := range *lib {
		book := &(*lib)[i]
		book.IsBorrowed = !book.IsBorrowed
	}
}
