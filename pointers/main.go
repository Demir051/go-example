package main

import "fmt"

func main() {

	/*
		var a int = 10
		var b *int = &a
		fmt.Println(b)
		fmt.Println(*b)
	*/

	/*
		var a int = 10
		add12(a)
		fmt.Println("Non pointer: ", a)
		add12Pointer(&a)
		fmt.Println("Pointer: ", a)
	*/

	var numbers = []int{1, 2, 3, 4, 5}
	change(numbers)
	fmt.Println(numbers)

}

func change(numbers []int) {
	numbers[0] = 99
}

func add12(x int) {
	x = x + 12
}

func add12Pointer(x *int) {
	*x = *x + 12
}
