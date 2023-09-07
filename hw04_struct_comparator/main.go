package main

import (
	"fmt"
	"os"
)

type BookComparerMode string

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}
type BookComparer struct {
	Mode string
}

func main() {
	book1 := NewBook(1, "1984", "Оруэлл Дж.", 1945, 352, 4.6)
	book2 := NewBook(2, "Animal Farm: A Fairy Story", "George Orwell", 1949, 986, 4.0)

	var comparerMode string
	fmt.Printf("Введите один из критериев для сравнения:\nByYear, BySize, ByRate\n")
	fmt.Fscan(os.Stdin, &comparerMode)
	comparer := &BookComparer{Mode: comparerMode}

	if comparer.Compare(book1, book2) {
		fmt.Printf("%s больше чем %s", book1.title, book2.title)
	} else {
		fmt.Printf("%s больше чем %s", book2.title, book1.title)
	}
}

func NewBook(id int, title, author string, year, size int, rate float64) *Book {
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

func (c *BookComparer) Compare(book1, book2 *Book) bool {
	switch c.Mode {
	case "ByYear":
		return book1.year > book2.year
	case "BySize":
		return book1.size > book2.size
	case "ByRate":
		return book1.rate > book2.rate
	default:
		return false
	}
}
