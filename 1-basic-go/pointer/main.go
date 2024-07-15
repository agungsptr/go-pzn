package main

import "fmt"

func main() {
	book1 := book{
		title:       "Outliers",
		author:      "Malcom Gladwell",
		publisher:   "Gramedia",
		isPublished: true,
	}
	fmt.Println("book1: ", book1)

	// Saat book1 di-assign ke book2 hanya valuenya yg diambil
	// sehigga saat book2 diubah maka tidak akan mengubah book2
	book2 := book1
	book2.isPublished = false
	fmt.Println("book1: ", book1)
	fmt.Println("book2: ", book2)

	// Jika menggunakan pointer maka value pada book1 akan berubah jika
	// mengubah book3
	// book 3,4,5 hanya alternatif cara pendeklarasian
	book3 := &book1
	var book4 *book = &book1
	book5 := new(book)
	book5 = &book1

	book3.isPublished = false

	fmt.Println("book1: ", book1)
	fmt.Println("book3: ", book3)
	fmt.Println("book4: ", book4)
	fmt.Println("book5: ", book5)
}

type book struct {
	title       string
	author      string
	publisher   string
	isPublished bool
}
