package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(result)
	} //gk bisa dikonversi jadi boolean
	// kalau true bisa

	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(resultInt)
	}

	binary := strconv.FormatInt(9999, 2)
	fmt.Println(binary)

	var stringINT string = strconv.Itoa((999))
	fmt.Println(stringINT)
}
