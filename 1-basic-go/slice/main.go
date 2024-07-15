package main

import "fmt"

func main() {
	// [] adalah tipe data slice, sedangkan [3] adalah array
	// slice datanya dapat ditambahkan, array tidak
	
	// 2 = size, 3 = capacity
	people := make([]string, 2, 3)
	people[0] = "Agung"
	people[1] = "Yuda"
	people = append(people, "Kuswari")
	fmt.Println(people)

	books := []string{"Outliers", "Tipping Point", "Funiculi Funicula"}
	fmt.Println(books)
}
