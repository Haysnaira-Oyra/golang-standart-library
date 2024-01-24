package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"Aryo", "Prasetio"})
	_ = writer.Write([]string{"Ariansyah", "Aryo"})
	_ = writer.Write([]string{"Prasetio", "Ariansyah"})

	writer.Flush()
}
