package main

import "fmt"

func main() {
	firstName := "Ariansyah"
	lastName := "Aryo"

	fmt.Println("Hello", firstName, lastName)
	fmt.Println("Hello'", firstName, lastName, "'")
	//	disubtitusikan biar gk ada jarak spasinya
	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
	fmt.Println("Hello", firstName, lastName, "!")
}
