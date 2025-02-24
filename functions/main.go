package main

import "fmt"

func main() {

	/*result := add(10, 45)
	fmt.Println(result)
	toal, diff := calculate(10, 45)
	fmt.Println(toal, diff) */

	/*var numbers = []int{1, 2, 3, 4, 5}
	sliced := sliceNumbers(numbers)
	fmt.Println(sliced) */
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func sum(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func sliceNumbers(x []int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	return sum
}

func calculate(a int, b int) (int, int) {
	return a + b, a - b
}
func add(a int, b int) int {
	return a + b
}
