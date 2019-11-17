package main

import "fmt"

type Book struct {
	name string
	number int
	author string
}

func main() {
	var book Book
	book.name = "go语言"
	book.number = 150
	book.author = "海明威"
	fmt.Print(book)
}
