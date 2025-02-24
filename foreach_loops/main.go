package main

import (
	"fmt"
	"strings"
)

func main() {

	/*var numbers = []int{1, 2, 3, 4, 5}

	for _, value := range numbers {
		fmt.Println(value)
	} */

	var name = "Demolingo"
	for _, value := range name {
		new := strings.ToUpper(string(value))
		fmt.Println(new)
	}
}
