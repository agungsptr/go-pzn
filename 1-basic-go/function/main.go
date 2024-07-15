package main

import "fmt"

func main() {
	firstName, lastName := getName()
	fmt.Printf("%s %s\n", firstName, lastName)

	nameFirst, nameMid, nameLast := getFullName()
	fmt.Printf("%s %s %s\n", nameFirst, nameMid, nameLast)

	// Func as Value
	calc := sum
	fmt.Println(calc(10, 10, 10))

	// Func as Parameter
	sumWithPrefix("Total:", sum, 1, 2, 3)
}

// Func Return Multiple Value
func getName() (string, string) {
	firstName := "Agung"
	lastName := "Saputra"

	return firstName, lastName
}

// Func Return (Named) Multiple Value
func getFullName() (firstName, middleName, lastName string) {
	firstName = "Agung"
	middleName = "Eka"
	lastName = "saputra"

	return firstName, middleName, lastName
}

// Type Function
type SumCalculator func(...int8) int8

// Func Variadic
func sum(nums ...int8) int8 {
	var total int8
	for _, num := range nums {
		total += num
	}
	return total
}

// Func as Parameter
func sumWithPrefix(prefix string, calc SumCalculator, nums ...int8) {
	fmt.Println(prefix, calc(nums...))
}
