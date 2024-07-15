package main

import "fmt"

func main() {
	book := map[string]string{
		"title":         "Ouliers",
		"author":        "Malcom Gladwell",
		"publisher":     "Gramedia",
		"publishStatus": "Published",
	}

	// IF Condition
	if book["publishStatus"] == "Published" {
		fmt.Printf("Book with title %s is published\n", book["title"])
	}

	// IF Shorthand Condition
	if publishStatus := book["publishStatus"]; publishStatus == "Published" {
		fmt.Printf("Book with title %s is published\n", book["title"])
	} else {
		fmt.Printf("Book with title %s is not published\n", book["title"])
	}

	// Switch Shorthand Condition
	switch publisher := book["publisher"]; publisher {
	case "Gramedia":
		fmt.Println("Book published by Gramedia")
	case "Intan Pariwara":
		fmt.Println("Book published by Intan Pariwara")
	default:
		fmt.Println("Book don't have publisher")
	}
}
