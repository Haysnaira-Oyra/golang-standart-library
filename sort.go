package main

import (
	"fmt"
	"sort"
)

type user struct {
	name string
	age  int
}

type UserSlice []user

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]

	//menukar posisi dengan temporari
	//temp := s[i]
	//s[i] = s[j]
	//s[j] = temp
}

func main() {
	users := []user{
		{"Aryo", 20},
		{"Prasetio", 21},
		{"Ariansyah", 22},
		{"Ary", 23},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
