package main

import "fmt"

// Define a struct Book with fields Title, Author, and Price.
//  Create a slice of Book structs and
// write a function to display all books priced below a given value.

type Book struct {
	Title  string
	Author string
	Price  float64
}

func filterBook(books []Book, price float64) []string {
	var filteredBooks []string
	for _, book := range books {
		if book.Price < price {
			filteredBooks = append(filteredBooks, book.Title)
		}
	}
	return filteredBooks
}

func main() {
	books := []Book{
		{Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Price: 40.00},
		{Title: "Clean Code", Author: "Robert C. Martin", Price: 35.50},
		{Title: "Introduction to Algorithms", Author: "Thomas H. Cormen", Price: 75.00},
		{Title: "The Pragmatic Programmer", Author: "Andrew Hunt", Price: 30.00},
	}
	booksBelow50 := filterBook(books, 50)
	fmt.Println(booksBelow50)
}
