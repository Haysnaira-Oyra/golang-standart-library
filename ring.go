package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}
	//data.Value = "value1"
	//
	//data = data.Next()
	//data.Value = "Value2"
	//
	//data = data.Next()
	//data.Value = "Value3"
	//
	//data = data.Next()
	//data.Value = "Value4"
	//
	//data = data.Next()
	//data.Value = "Value5"

	data.Do(func(value any) {
		fmt.Println(value)
	})

}
