package main

import "fmt"

func main() {
	book := book{
		title:       "Outliers",
		author:      "Malcom Gladwell",
		publisher:   "Gramedia",
		isPublished: true,
	}
	fmt.Println(book)
}

type book struct {
	title       string
	author      string
	publisher   string
	isPublished bool
}
