package main

import (
	"fmt"
)

type Book struct {
	id     int
	name   string
	author string
	tag    []string
	price  float64
}

func (b *Book) SetName(name string) {
	b.name = name
}

func main() {
	b := Book{
		1,
		"一千零一夜",
		"阿拉伯",
		[]string{"故事", "寓言", "传说", "奇幻"},
		50.00,
	}
	fmt.Println("book:", b)

	var b1 Book
	b1.id = 2
	b1.name = "伊索寓言"

	fmt.Println("book1:", b1)

	var b2 *Book
	b2 = new(Book)

	fmt.Println("book2:", *b2)
}

// var b Book
