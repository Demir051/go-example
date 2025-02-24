package main

import "fmt"

func main() {

	index := 1

	for index <= 10 {
		println(index)
		index++
	}

	total := 0
	for index := 1; index <= 10; index++ {
		//total = index + total
		total += index
	}
	fmt.Println(total)

	i := 0
	var names = [3]string{"John", "Ahmet", "Demir"}
	for i < len(names) {
		if i == 2 {
			break
		} else if i == 1 {
			continue
		}
		fmt.Println(names[i])
		i++
	}

}
