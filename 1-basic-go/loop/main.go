package main

import "fmt"

func main() {
	// For Loop
	for i := 0; i < 3; i++ {
		fmt.Printf("Ping: %d\n", i+1)
	}

	// For Range Loop
	books := [...]string{"Outliers", "Tipping Point", "Funiculi Funicula"}
	for i, book := range books {
		fmt.Printf("Book-%d: %s\n", i+1, book)
	}

	for _, book := range books {
		fmt.Printf("Book with title: %s\n", book)
	}
}
