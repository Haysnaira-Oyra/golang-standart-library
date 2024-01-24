package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("Golang-Standart-Library/path.go"))
	fmt.Println(path.Base("Golang-Standart-Library/path.go"))
	fmt.Println(path.Ext("/home/aryo/go/src/github.com/aryo13/learn-golang/Golang-Standart-Library/path.go"))
	fmt.Println(path.Join("/home/aryo/go/src/github.com/aryo13/learn-golang/Golang-Standart-Library", "path.go"))
}
