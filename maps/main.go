package main

import "fmt"

func main() {

	/*var names = map[string]string{
		"names": "Doe",
		"age":   "21",
	} */

	/*names := make(map[string]string)
	names["names"] = "Doe"
	names["age"] = "21"
	fmt.Println(names) */

	names := map[string]string{
		"name": "Doe",
		"age":  "21",
	}
	delete(names, "age")
	fmt.Println(names)
}
