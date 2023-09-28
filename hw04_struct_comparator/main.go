package main

import (
	"fmt"
)

// Book - структура для представления информации о книге
type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

// NewBook - конструктор для создания новой книги
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

// Сеттеры для возможности утсновки значения поля
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

// Геттеры для возможности получения значения поля

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

// BookComparerMode - перечисление для выбора режима сравнения
type BookComparerMode int

const (
	ByYear BookComparerMode = iota
	BySize
	ByRate
)

// BookComparer - структура для сравнения книг по выбранному полю
type BookComparer struct {
	Mode BookComparerMode
}

// Compare - метод для сравнения двух книг по выбранному полю
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
	// Создаем объект BookComparer с выбранным режимом сравнения
	comparer := &BookComparer{Mode: BySize}

	// Сравниваем книги по выбранному полю
	if comparer.Compare(book1, book2) {
		fmt.Println("Book 1 is greater than Book 2")
	} else {
		fmt.Println("Book 2 is greater than Book 1")
	}
}
