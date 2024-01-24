package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Aryo", "Prasetio", "Ariansyah"}
	values := []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Aryo"))
	fmt.Println(slices.Contains(names, "Arya"))
	fmt.Println(slices.Contains(values, 1))
}
