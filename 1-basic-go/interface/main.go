package main

import "fmt"

func main() {
	// Menggunakan interface sebagai tipe variable,
	// dan meng-assign var person dg struct Person
	// yg mengimplement interface HasName
	var person HasName = Person{Name: "Agung Saputra"}

	fmt.Println(person.getName())
}

// Interface HasName
type HasName interface {
	getName() string
}

// Implement interface dg membuat struct mengikuti kontrak interface.
// Interface HasName memiliki function getName, maka dg membuat
// struct yg memilki function getName artinya sudah mengimplement
// interface HasName
type Person struct {
	Name string
}

// Function dr struct Person, ditandai dg membawa nilai struct itu sendiri.
func (p Person) getName() string {
	return p.Name
}
