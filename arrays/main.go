package main

import "fmt"

func main() {
	//var names [3]string
	//names[0] = "John"
	//names[1] = "Ahmet"
	//names[2] = "Demir"

	var names = []string{"John", "Ahmet", "Demir"}
	names[2] = "Kızgın"
	names = append(names, "Kuzgun")
	fmt.Println(names)
	fmt.Println(names[0:2])
}
