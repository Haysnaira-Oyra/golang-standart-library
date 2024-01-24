package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("Golang-Standart-Library/file_path.go"))
	fmt.Println(filepath.Base("Golang-Standart-Library/file_path.go"))
	fmt.Println(filepath.Ext("/home/aryo/go/src/github.com/aryo13/learn-golang/Golang-Standart-Library/file_path.go"))
	fmt.Println(filepath.IsAbs("/home/aryo/go/src/github.com/aryo13/learn-golang/Golang-Standart-Library/file_path.go"))
	//fmt.Println(filepath.IsLocal("/home/aryo/go/src/github.com/aryo13/learn-golang/Golang-Standart-Library/file_path.go"))
	fmt.Println(filepath.Join("/home/aryo/go/src/github.com/aryo13/learn-golang/Golang-Standart-Library", "file_path.go"))

}
