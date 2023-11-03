package main

import (
	"fmt"
)

func main() {

	// variables && types && pointers
	var name string = "Demir"
	var namePointer *string = &name
	fmt.Println(name, namePointer, *namePointer)

	fmt.Println(tr, usd, euro, gbp)

	// arrays
	arr := [...]int{1, 2, 3, 4, 5}
	slice := arr[1:3]
	fmt.Println("len =>", len(arr), "cap =>", cap(arr), "sliece =>", slice)

	// maps
	m := map[string]int{"istanbul": 1}
	m["ankara"] = 2
	fmt.Println(m, "/////", m["istanbul"], m["ankara"])

}

// iota
const (
	tr = iota + 1
	usd
	euro
	gbp
)
