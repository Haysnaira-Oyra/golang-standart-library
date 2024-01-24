package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])o")
	fmt.Println(regex.MatchString("Ayo belajar"))
	fmt.Println(regex.MatchString("ayo"))
	fmt.Println(regex.MatchString("aryi"))
	fmt.Println(regex.MatchString("eryo"))
	fmt.Println(regex.MatchString("ato"))

	fmt.Println(regex.FindAllString("ayo aryi aryu eryo oryo arye eryo orya ato", 10))

}
