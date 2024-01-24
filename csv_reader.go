package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Ariansyah,Aryo,Prasetio\n" +
		"Ryo,Arya,Prasetyo\n" +
		"Joko,Anwar,Prasetyo"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
