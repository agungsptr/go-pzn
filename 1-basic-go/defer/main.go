package main

import "fmt"

func main() {
	app()
}

func logger() {
	fmt.Println("App stopped")
}

func app() {
	// Defer (run after end execution of function)
	defer logger()
	fmt.Println("Running app...")
}
