package main

import (
	"fmt"
	"strings"
)

// Ini buat manipulasi string
func main() {
	fmt.Println(strings.Contains("Aryo Prasetio", "Aryo"))
	fmt.Println(strings.Split("Aryo Prasetio", " "))
	fmt.Println(strings.ToLower("Ariansyah Aryo"))
	fmt.Println(strings.ToUpper("Aryo Prasetio"))
	fmt.Println(strings.Trim("	Aryo Prasetio	", ""))
	fmt.Println(strings.ReplaceAll("Aryo Prasetio Aryo Ariansyah", "Aryo", "Ryo"))
}
