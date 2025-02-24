package main

func main() {

	a := 10
	b := 40
	c := 30

	if a > b && a > c {
		println("a is the largest")
	} else if b > a && b > c {
		println("b is the largest")
	} else {
		println("c is the largest")
	}
}
