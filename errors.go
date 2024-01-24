package main

import (
	"errors"
	"fmt"
)

// mau var atau const terserah
// tapi kalau const ada minesnya yaitu gk bisa ngembalikan function
var (
	ValidationError = errors.New("Validasi Error")
	NotFoundError   = errors.New("Error Tidak ditemukan")
)

func GetByID(id string) error {
	if id == "" {
		return ValidationError
	}
	if id != "Aryo" {
		return NotFoundError
	}
	// sukses tidak terjadi error
	return nil
}

func main() {
	err := GetByID("Aryo")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validasi Error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("Unknown Error")
		}
	}
}
