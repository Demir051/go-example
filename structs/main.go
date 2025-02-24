package main

import "fmt"

func main() {

	customer1 := Customer{
		name: "Demir",
		age:  20,
		adress: Adress{
			street: "Cumhuriyet",
			city:   "Kocaeli",
		},
	}
	customer1.setName("Ahmet")
	customer1.toString()
}

type Customer struct {
	name   string
	age    int
	adress Adress
}

type Adress struct {
	street string
	city   string
}

func (customer Customer) toString() {
	fmt.Printf("Name: %s\nAge: %d\nStreet: %s\nCity: %s\n", customer.name, customer.age, customer.adress.street, customer.adress.city)
}

func (customer *Customer) setName(name string) {
	customer.name = name
}
