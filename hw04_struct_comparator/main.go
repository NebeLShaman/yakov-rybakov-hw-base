package main

import (
	"fmt"
)

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
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

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float64) {
	b.rate = rate
}

func (b *Book) GetID() int {
	return b.id
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) GetYear() int {
	return b.year
}

func (b *Book) GetSize() int {
	return b.size
}

func (b *Book) GetRate() float64 {
	return b.rate
}

type BookComparerMode int

const (
	ByYear BookComparerMode = iota
	BySize
	ByRate
)

type BookComparer struct {
	Mode BookComparerMode
}

func (c *BookComparer) Compare(book1, book2 *Book) bool {
	switch c.Mode {
	case ByYear:
		return book1.year < book2.year
	case BySize:
		return book1.size > book2.size
	case ByRate:
		return book1.rate > book2.rate
	default:
		return false
	}
}

func main() {
	book1 := NewBook(1, "1984", "Оруэлл Дж.", 1945, 352, 4.6)
	book2 := NewBook(2, "Animal Farm: A Fairy Story", "George Orwell", 1949, 986, 4.0)
	comparer := &BookComparer{Mode: BySize}

	if comparer.Compare(book1, book2) {
		fmt.Println("Book 1 is greater than Book 2")
	} else {
		fmt.Println("Book 2 is greater than Book 1")
	}
}
