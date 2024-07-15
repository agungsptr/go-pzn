package main

import "fmt"

func main() {
	app(true)
}

// Panic and Recover (to stop running program)
// Try - Catch in Go
func logError() {
	fmt.Println("Program has stop")

	// Recover harus di panggil di defer fucntion,
	// karena hanya defer yg dijalankan terkahir meski program stop
	errorMsg := recover()
	fmt.Println("Error:", errorMsg)
}

func app(error bool) {
	if error {
		panic("Error signal appear")
	}
	defer logError()

	fmt.Println("App running...")
}
