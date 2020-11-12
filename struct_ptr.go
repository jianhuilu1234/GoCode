package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	var Book1 Books
	var Book2 Books

	Book1.title = "Go 语言"
	Book1.author = "lujh"
	Book1.subject = "教程"
	Book1.bookId = 1

	Book2.title = "Python 语言"
	Book2.author = "lujh"
	Book2.subject = "教程"
	Book2.bookId = 2

	printBook_by_ptr(&Book1)
	printBook_by_ptr(&Book2)
}
func printBook_by_ptr(book *Books) {
	fmt.Println(book.bookId)
	fmt.Println(book.subject)
	fmt.Println(book.author)
	fmt.Println(book.title)
}
